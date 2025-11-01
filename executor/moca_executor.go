package executor

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"time"

	"cosmossdk.io/math"
	"github.com/0xPolygon/polygon-edge/bls"
	"github.com/avast/retry-go/v4"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/cometbft/cometbft/votepool"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	oracletypes "github.com/cosmos/cosmos-sdk/x/oracle/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"

	gnfdsdktypes "github.com/evmos/evmos/v12/sdk/types"
	sdktypes "github.com/mocachain/moca-go-sdk/types"
	relayercommon "github.com/mocachain/moca-relayer/common"
	"github.com/mocachain/moca-relayer/config"
	"github.com/mocachain/moca-relayer/contract/universalVerifier"
	"github.com/mocachain/moca-relayer/contract/mocacrosschainupgradeable"
	"github.com/mocachain/moca-relayer/logging"
	"github.com/mocachain/moca-relayer/types"
)

type MocaExecutor struct {
	mutex         sync.RWMutex
	BscExecutor   *BSCExecutor
	gnfdClients   GnfdCompositeClients
	config        *config.Config
	privateKey    *ecdsa.PrivateKey
	address       string
	validators    []*tmtypes.Validator // used to cache validators
	BlsPrivateKey []byte
	BlsPubKey     []byte
}

func NewMocaExecutor(cfg *config.Config) *MocaExecutor {
	privKey := viper.GetString(config.FlagConfigPrivateKey)
	if privKey == "" {
		privKey = getMocaPrivateKey(&cfg.MocaConfig)
	}
	ecdsaPrivKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		panic(err)
	}
	blsPrivKeyStr := viper.GetString(config.FlagConfigBlsPrivateKey)
	if blsPrivKeyStr == "" {
		blsPrivKeyStr = getMocaBlsPrivateKey(&cfg.MocaConfig)
	}
	blsPrivKeyBts, err := hex.DecodeString(blsPrivKeyStr)
	if err != nil {
		panic(err)
	}
	blsPrivKey, err := bls.UnmarshalPrivateKey(blsPrivKeyBts)
	if err != nil {
		panic(err)
	}
	account, err := sdktypes.NewAccountFromPrivateKey("relayer", privKey)
	if err != nil {
		panic(err)
	}
	clients := NewGnfdCompositClients(
		cfg.MocaConfig.RPCAddrs,
		cfg.MocaConfig.EVMAddrs,
		cfg.MocaConfig.ChainIdString,
		cfg.MocaConfig.PrivateKey,
		account,
		cfg.MocaConfig.UseWebsocket,
		cfg.RelayConfig.SrcMocaSBTContractAddr,
		cfg.RelayConfig.SrcMocaVCContractAddr,
	)
	return &MocaExecutor{
		gnfdClients:   clients,
		address:       account.GetAddress().String(),
		config:        cfg,
		privateKey:    ecdsaPrivKey,
		BlsPrivateKey: blsPrivKeyBts,
		BlsPubKey:     blsPrivKey.PublicKey().Marshal(),
	}
}

func (e *MocaExecutor) SetBSCExecutor(be *BSCExecutor) {
	e.BscExecutor = be
}

func getMocaPrivateKey(cfg *config.MocaConfig) string {
	if cfg.KeyType == config.KeyTypeAWSPrivateKey {
		result, err := config.GetSecret(cfg.AWSSecretName, cfg.AWSRegion)
		if err != nil {
			panic(err)
		}
		type AwsPrivateKey struct {
			PrivateKey string `json:"private_key"`
		}
		var awsPrivateKey AwsPrivateKey
		err = json.Unmarshal([]byte(result), &awsPrivateKey)
		if err != nil {
			panic(err)
		}
		return awsPrivateKey.PrivateKey
	}
	return cfg.PrivateKey
}

func getMocaBlsPrivateKey(cfg *config.MocaConfig) string {
	if cfg.KeyType == config.KeyTypeAWSPrivateKey {
		result, err := config.GetSecret(cfg.AWSBlsSecretName, cfg.AWSRegion)
		if err != nil {
			panic(err)
		}
		type AwsPrivateKey struct {
			PrivateKey string `json:"bls_private_key"`
		}
		var awsBlsPrivateKey AwsPrivateKey
		err = json.Unmarshal([]byte(result), &awsBlsPrivateKey)
		if err != nil {
			panic(err)
		}
		return awsBlsPrivateKey.PrivateKey
	}
	return cfg.BlsPrivateKey
}

func (e *MocaExecutor) GetGnfdClient() *MocaClient {
	return e.gnfdClients.GetClient()
}

func (e *MocaExecutor) GetEthClient() *ethclient.Client {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.gnfdClients.GetClient().ethClient
}

func (e *MocaExecutor) getSBTCrossChainClient() *mocacrosschainupgradeable.IMOCACrossChainUpgradeable {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.gnfdClients.GetClient().mocaSBTCrossChainClient
}

func (e *MocaExecutor) getVCCrossChainClient() *universalVerifier.UniversalVerifier {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.gnfdClients.GetClient().mocaVCCrossChainClient
}

func (e *MocaExecutor) GetBlockAndBlockResultAtHeight(height int64) (*tmtypes.Block, *ctypes.ResultBlockResults, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	block, err := e.GetGnfdClient().GetBlockByHeight(ctx, height)
	if err != nil {
		return nil, nil, err
	}
	blockResults, err := e.GetGnfdClient().GetBlockResultByHeight(ctx, height)
	if err != nil {
		return nil, nil, err
	}
	return block, blockResults, nil
}

func (e *MocaExecutor) GetLatestBlockHeight() (latestHeight uint64, err error) {
	return uint64(e.gnfdClients.GetClient().Height), nil
}

func (e *MocaExecutor) QueryTendermintLightBlock(height int64) (tmtypes.LightBlock, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	validators, err := e.GetGnfdClient().GetValidatorsByHeight(ctx, height)
	if err != nil {
		return tmtypes.LightBlock{}, err
	}
	commit, err := e.GetGnfdClient().GetCommit(ctx, height)
	if err != nil {
		return tmtypes.LightBlock{}, err
	}
	validatorSet := tmtypes.NewValidatorSet(validators)
	if err != nil {
		return tmtypes.LightBlock{}, err
	}
	lightBlock := tmtypes.LightBlock{
		SignedHeader: &commit.SignedHeader,
		ValidatorSet: validatorSet,
	}
	return lightBlock, nil
}

// GetNextDeliverySequenceForChannelWithRetry calls dest chain(BSC) to return a sequence # which should be used.
func (e *MocaExecutor) GetNextDeliverySequenceForChannelWithRetry(channelID types.ChannelId) (sequence uint64, err error) {
	return sequence, retry.Do(func() error {
		sequence, err = e.getNextDeliverySequenceForChannel(channelID)
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			logging.Logger.Errorf("failed to query receive sequence for channel %d, attempt: %d times, max_attempts: %d", channelID, n+1, relayercommon.RtyAttNum)
		}))
}

func (e *MocaExecutor) getNextDeliverySequenceForChannel(channelID types.ChannelId) (uint64, error) {
	sequence, err := e.BscExecutor.GetNextReceiveSequenceForChannelWithRetry(channelID)
	if err != nil {
		return 0, err
	}
	return sequence, nil
}

// GetNextSendSequenceForChannelWithRetry gets the next send sequence of a specified channel from Moca
func (e *MocaExecutor) GetNextSendSequenceForChannelWithRetry(destChainID sdk.ChainID, channelID types.ChannelId) (sequence uint64, err error) {
	return sequence, retry.Do(func() error {
		if channelID == relayercommon.MocaSBTChannelId {
			sequence, err = e.getNextMocaSBTSendSequenceForChain(destChainID)
		} else if channelID == relayercommon.MocaVCChannelId {
			sequence, err = e.getNextMocaVCSendSequenceForChain(destChainID)
		} else {
			sequence, err = e.getNextSendSequenceForChannel(destChainID, channelID)
		}
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			logging.Logger.Errorf("failed to query send sequence for chain %d, channel %d, attempt: %d times, max_attempts: %d", destChainID, channelID, n+1, relayercommon.RtyAttNum)
		}))
}

func (e *MocaExecutor) getNextSendSequenceForChannel(destChainId sdk.ChainID, channelId types.ChannelId) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	return e.GetGnfdClient().GetChannelSendSequence(
		ctx,
		destChainId,
		uint32(channelId),
	)
}

func (e *MocaExecutor) getNextMocaSBTSendSequenceForChain(destChainId sdk.ChainID) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	callOpts := &bind.CallOpts{
		Pending: true,
		Context: ctx,
	}
	seq, err := e.getSBTCrossChainClient().GetCrossChainSequence(callOpts, uint32(destChainId))
	if err != nil {
		return 0, err
	}
	return seq.Uint64(), err
}

func (e *MocaExecutor) getNextMocaVCSendSequenceForChain(destChainId sdk.ChainID) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	callOpts := &bind.CallOpts{
		Pending: true,
		Context: ctx,
	}
	seq, err := e.getVCCrossChainClient().GetVCCrossChainSequence(callOpts, uint32(destChainId))
	if err != nil {
		return 0, err
	}
	return seq.Uint64(), err
}

// GetNextReceiveOracleSequence gets the next receive Oracle sequence from Moca
func (e *MocaExecutor) GetNextReceiveOracleSequence(destChainId sdk.ChainID) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	return e.GetGnfdClient().GetChannelReceiveSequence(
		ctx,
		destChainId,
		uint32(relayercommon.OracleChannelId),
	)
}

// GetNextReceiveSequenceForChannel gets the sequence specifically for bsc -> gnfd package's channel from Moca
func (e *MocaExecutor) GetNextReceiveSequenceForChannel(destChainId sdk.ChainID, channelId types.ChannelId) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	return e.GetGnfdClient().GetChannelReceiveSequence(
		ctx,
		destChainId,
		uint32(channelId),
	)
}

func (e *MocaExecutor) queryLatestValidators() ([]*tmtypes.Validator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	_, validators, err := e.GetGnfdClient().GetValidatorSet(ctx)
	if err != nil {
		return nil, err
	}
	return validators, nil
}

func (e *MocaExecutor) QueryValidatorsAtHeight(height uint64) ([]*tmtypes.Validator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	return e.GetGnfdClient().GetValidatorsByHeight(ctx, int64(height))
}

func (e *MocaExecutor) QueryCachedLatestValidators() ([]*tmtypes.Validator, error) {
	if len(e.validators) != 0 {
		return e.validators, nil
	}
	validators, err := e.queryLatestValidators()
	if err != nil {
		return nil, err
	}
	return validators, nil
}

func (e *MocaExecutor) UpdateCachedLatestValidatorsLoop() {
	ticker := time.NewTicker(UpdateCachedValidatorsInterval)
	for range ticker.C {
		validators, err := e.queryLatestValidators()
		if err != nil {
			logging.Logger.Errorf("update latest moca validators error, err=%s", err)
			continue
		}
		e.validators = validators
	}
}

func (e *MocaExecutor) GetValidatorsBlsPublicKey() ([]string, error) {
	validators, err := e.QueryCachedLatestValidators()
	if err != nil {
		return nil, err
	}
	var keys []string
	for _, v := range validators {
		keys = append(keys, hex.EncodeToString(v.BlsKey))
	}
	return keys, nil
}

func (e *MocaExecutor) GetNonce() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	acc, err := e.GetGnfdClient().GetAccount(ctx, e.address)
	if err != nil {
		return 0, err
	}
	return acc.GetSequence(), nil
}

func (e *MocaExecutor) GetNonceOnNextBlock() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	err := e.GetGnfdClient().WaitForNextBlock(ctx)
	if err != nil {
		return 0, err
	}
	return e.GetNonce()
}

func (e *MocaExecutor) getGasPrice() (*big.Int, error) {
	var (
		gasPrice *big.Int
		err      error
	)
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	if e.config.BSCConfig.GasPrice == 0 {
		gasPrice, err = e.GetEthClient().SuggestGasPrice(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		gasPrice = big.NewInt(int64(e.config.BSCConfig.GasPrice))
	}
	return gasPrice, nil
}

// TODO
func (e *MocaExecutor) getTransactor(nonce uint64) (*bind.TransactOpts, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(e.privateKey, big.NewInt(int64(e.config.MocaConfig.ChainId)))
	if err != nil {
		return nil, err
	}
	gasPrice, err := e.getGasPrice()
	if err != nil {
		return nil, err
	}
	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.Value = big.NewInt(0)
	txOpts.GasLimit = e.config.BSCConfig.GasLimit
	txOpts.GasPrice = big.NewInt(gasPrice.Int64() + 1)
	return txOpts, nil
}

func (e *MocaExecutor) CallMocaSBTAckMintedContract(chainId uint32, user ethcommon.Address, status uint8, nonce uint64) (ethcommon.Hash, error) {
	txOpts, err := e.getTransactor(nonce)
	if err != nil {
		return ethcommon.Hash{}, err
	}

	tx, err := e.getSBTCrossChainClient().AckMinted(txOpts, chainId, user, status)
	if err != nil {
		return ethcommon.Hash{}, err
	}
	return tx.Hash(), nil
}

func (e *MocaExecutor) ClaimPackages(client *MocaClient, payloadBts []byte, aggregatedSig []byte, voteAddressSet []uint64, claimTs int64, oracleSeq uint64, nonce uint64) (string, error) {
	msg := oracletypes.NewMsgClaim(
		e.address,
		e.getSrcChainId(),
		e.getDestChainId(),
		oracleSeq,
		uint64(claimTs),
		payloadBts,
		voteAddressSet,
		aggregatedSig)
	gasLimit, feeAmount, err := e.getGasLimitAndFeeAmount(msg)
	if err != nil {
		return "", err
	}
	txOpt := gnfdsdktypes.TxOption{
		NoSimulate: true,
		GasLimit:   uint64(gasLimit),
		FeeAmount:  sdk.NewCoins(sdk.NewCoin(gnfdsdktypes.Denom, math.NewInt(feeAmount))),
		Nonce:      nonce,
	}
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	resp, err := client.BroadcastTx(ctx, []sdk.Msg{msg}, &txOpt)
	if err != nil {
		return "", err
	}
	txRes := resp.TxResponse
	if txRes.Codespace == oracletypes.ModuleName && txRes.Code == oracletypes.ErrInvalidReceiveSequence.ABCICode() {
		return "", oracletypes.ErrInvalidReceiveSequence
	}

	if txRes.Codespace == sdkErrors.RootCodespace && txRes.Code == sdkErrors.ErrWrongSequence.ABCICode() {
		return "", sdkErrors.ErrWrongSequence
	}

	if txRes.Codespace == sdkErrors.RootCodespace && txRes.Code == sdkErrors.ErrTxInMempoolCache.ABCICode() {
		return "", sdkErrors.ErrTxInMempoolCache
	}

	if txRes.Code != 0 {
		return "", fmt.Errorf("claim error, code=%d, log=%s", txRes.Code, txRes.RawLog)
	}
	return txRes.TxHash, nil
}

func (e *MocaExecutor) GetInturnRelayer(srcChain oracletypes.ClaimSrcChain) (*oracletypes.QueryInturnRelayerResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	return e.GetGnfdClient().GetInturnRelayer(ctx, &oracletypes.QueryInturnRelayerRequest{
		ClaimSrcChain: srcChain,
	})
}

func (e *MocaExecutor) QueryVotesByEventHashAndType(eventHash []byte, eventType votepool.EventType) ([]*votepool.Vote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	votes, err := e.gnfdClients.GetClient().QueryVote(ctx, int(eventType), eventHash)
	if err != nil {
		return nil, err
	}
	return votes.Votes, nil
}

func (e *MocaExecutor) BroadcastVote(v *votepool.Vote) error {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	return e.gnfdClients.GetClient().BroadcastVote(ctx, *v)
}

func (e *MocaExecutor) getDestChainId() uint32 {
	return uint32(e.config.MocaConfig.ChainId)
}

func (e *MocaExecutor) getSrcChainId() uint32 {
	return uint32(e.config.BSCConfig.ChainId)
}

func (e *MocaExecutor) getGasLimitAndFeeAmount(msg *oracletypes.MsgClaim) (gasLimit int64, feeAmount int64, err error) {
	bz, err := msg.Marshal()
	if err != nil {
		return
	}
	if len(bz)+EstimatedTxExtraMetaSize >= MaxTxSizeForFixGasLimit {
		gasLimit = GasLimitRatio * int64(len(bz)+EstimatedTxExtraMetaSize)
		feeAmount = gasLimit * GnfdGasPrice
		return
	}
	return e.config.MocaConfig.GasLimit, e.config.MocaConfig.FeeAmount, nil
}

func (e *MocaExecutor) GetCrossTxPack(destChainID sdk.ChainID, channelID types.ChannelId, sequence uint64) (pack []byte, err error) {
	return pack, retry.Do(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
		defer cancel()
		pack, err = e.GetGnfdClient().GetCrossChainPackage(ctx, destChainID, uint32(channelID), sequence)
		return err
	}, relayercommon.RtyAttem,
		relayercommon.RtyDelay,
		relayercommon.RtyErr,
		retry.OnRetry(func(n uint, err error) {
			logging.Logger.Errorf("failed to query crosschain tx for channel %d, seq %d, attempt: %d times, max_attempts: %d", channelID, n+1, relayercommon.RtyAttNum)
		}))
}
