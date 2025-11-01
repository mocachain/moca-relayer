package listener

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/cometbft/cometbft/votepool"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/mocachain/moca-relayer/common"
	"github.com/mocachain/moca-relayer/config"
	"github.com/mocachain/moca-relayer/contract/mocacrosschainupgradeable"
	"github.com/mocachain/moca-relayer/contract/universalVerifier"
	"github.com/mocachain/moca-relayer/db"
	"github.com/mocachain/moca-relayer/db/dao"
	"github.com/mocachain/moca-relayer/db/model"
	"github.com/mocachain/moca-relayer/executor"
	"github.com/mocachain/moca-relayer/logging"
	"github.com/mocachain/moca-relayer/metric"

	"github.com/mocachain/moca-relayer/types"
	"github.com/mocachain/moca-relayer/util"
)

type MocaListener struct {
	config           *config.Config
	mocaExecutor     *executor.MocaExecutor
	bscExecutor      *executor.BSCExecutor
	DaoManager       *dao.DaoManager
	sbtCrossChainAbi abi.ABI
	vcCrossChainAbi  abi.ABI
	metricService    *metric.MetricService
}

func NewMocaListener(cfg *config.Config, gnfdExecutor *executor.MocaExecutor, bscExecutor *executor.BSCExecutor,
	dao *dao.DaoManager, ms *metric.MetricService,
) *MocaListener {
	sbtCrossChainAbi, err := abi.JSON(strings.NewReader(mocacrosschainupgradeable.IMOCACrossChainUpgradeableMetaData.ABI))
	if err != nil {
		panic("marshal abi error")
	}
	vcCrossChainAbi, err := abi.JSON(strings.NewReader(universalVerifier.UniversalVerifierMetaData.ABI))
	if err != nil {
		panic("marshal abi error")
	}
	return &MocaListener{
		config:           cfg,
		mocaExecutor:     gnfdExecutor,
		bscExecutor:      bscExecutor,
		DaoManager:       dao,
		sbtCrossChainAbi: sbtCrossChainAbi,
		vcCrossChainAbi:  vcCrossChainAbi,
		metricService:    ms,
	}
}

func (l *MocaListener) StartLoop() {
	for {
		if err := l.poll(); err != nil {
			logging.Logger.Errorf("encounter err, err=%s", err.Error())
			time.Sleep(common.ErrorRetryInterval)
			continue
		}
	}
}

func (l *MocaListener) poll() error {
	nextHeight, err := l.calNextHeight()
	if err != nil {
		return fmt.Errorf("failed to cal next height, error: %s", err.Error())
	}
	blockResults, block, err := l.getBlockAndBlockResult(nextHeight)
	if err != nil {
		return fmt.Errorf("failed to get block and block result at height %d, error: %s", nextHeight, err.Error())
	}
	txs := make([]*model.MocaRelayTransaction, 0)
	wg := new(sync.WaitGroup)
	wg.Add(3)
	relayTxCh := make(chan *model.MocaRelayTransaction)
	errChan := make(chan error)
	waitCh := make(chan struct{})

	go func() {
		go l.monitorTxEvents(block, blockResults.TxsResults, relayTxCh, errChan, wg)
		go l.monitorEndBlockEvents(uint64(block.Height), blockResults.FinalizeBlockEvents, relayTxCh, errChan, wg)
		go l.monitorValidators(block, errChan, wg)
		wg.Wait()
		close(waitCh)
	}()

	for {
		select {
		case err := <-errChan:
			return fmt.Errorf("encounter error when monitoring block at Height=%d, err=%s", nextHeight, err.Error())
		case tx := <-relayTxCh:
			txs = append(txs, tx)
		case <-waitCh:
			// validate tx against Chain
			for _, tx := range txs {
				onchainPack, err := l.mocaExecutor.GetCrossTxPack(sdk.ChainID(tx.DestChainId), types.ChannelId(tx.ChannelId), tx.Sequence)
				if err != nil {
					return fmt.Errorf("failed to get on chain tx, err=%s", err.Error())
				}
				if err := l.validateTx(onchainPack, tx); err != nil {
					panic(fmt.Sprintf("failed to validate tx, tx.DestChainId=%d, tx.ChannelId=%d, tx.Sequence=%d. err=%s",
						tx.DestChainId, tx.ChannelId, tx.Sequence, err.Error()))
				}
			}

			b := &model.MocaBlock{
				Chain:     block.ChainID,
				Height:    uint64(block.Height),
				BlockTime: block.Time.Unix(),
			}

			if err := l.DaoManager.MocaDao.SaveBlockAndBatchTransactions(b, txs); err != nil {
				return fmt.Errorf("failed to persist block and tx to DB, err=%s", err.Error())
			}
			l.metricService.SetGnfdSavedBlockHeight(uint64(block.Height))
			return nil
		}
	}
}

func (l *MocaListener) validateTx(expectedPack []byte, tx *model.MocaRelayTransaction) error {
	relayerFee, err := util.StrToBigInt(tx.RelayerFee)
	if err != nil {
		return err
	}
	ackRelayerFee, err := util.StrToBigInt(tx.AckRelayerFee)
	if err != nil {
		return err
	}
	// package
	packBz := make([]byte, 0)

	// package header
	packageHeader := sdk.EncodePackageHeader(sdk.PackageHeader{
		PackageType:   sdk.CrossChainPackageType(tx.PackageType),
		Timestamp:     uint64(tx.TxTime),
		RelayerFee:    relayerFee,
		AckRelayerFee: ackRelayerFee,
	})
	packBz = append(packBz, packageHeader...)

	// package payload
	payloadBz, err := hex.DecodeString(tx.PayLoad)
	if err != nil {
		return err
	}
	packBz = append(packBz, payloadBz...)

	if !bytes.Equal(expectedPack, packBz) {
		return fmt.Errorf("package not match, expectedPack=%x, packBz=%x", expectedPack, packBz)
	}
	return nil
}

func (l *MocaListener) getLatestPolledBlock() (*model.MocaBlock, error) {
	return l.DaoManager.MocaDao.GetLatestBlock()
}

func (l *MocaListener) getBlockAndBlockResult(height uint64) (*ctypes.ResultBlockResults, *tmtypes.Block, error) {
	block, blockResults, err := l.mocaExecutor.GetBlockAndBlockResultAtHeight(int64(height))
	if err != nil {
		return nil, nil, err
	}
	logging.Logger.Infof("retrieved moca block at height=%d", height)
	return blockResults, block, nil
}

func (l *MocaListener) monitorTxEvents(block *tmtypes.Block, txRes []*abci.ExecTxResult, txChan chan *model.MocaRelayTransaction, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	// Cross chain Transfer events
	for idx, tx := range txRes {
		for _, event := range tx.Events {
			if event.Type == MocaEventTypeCrossChain {
				relayTx, err := constructRelayTx(event, uint64(block.Height))
				if err != nil {
					errChan <- err
					return
				}
				if relayTx.DestChainId != l.destChainId() {
					break
				}
				relayTx.TxHash = hex.EncodeToString(block.Txs[idx].Hash())
				txChan <- relayTx
			}
		}
	}

	logs, err := l.querySBTCrossChainLogs(uint64(block.Height))
	if err != nil {
		errChan <- err
		return
	}

	for _, log := range logs {
		logging.Logger.Infof("get log: %d, %s, %s", log.BlockNumber, log.Topics[0].String(), log.TxHash.String())
		relayTx, err := ParseMocaSBTRelayPackage(&l.sbtCrossChainAbi,
			&log, uint64(block.Header.Time.Unix()),
			sdk.ChainID(l.config.MocaConfig.ChainId),
			sdk.ChainID(l.config.BSCConfig.ChainId),
		)
		if err != nil {
			logging.Logger.Errorf("failed to parse event log, txHash=%s, err=%s", log.TxHash, err.Error())
			continue
		}
		txChan <- relayTx
	}

	logs, err = l.queryVCCrossChainLogs(uint64(block.Height))
	if err != nil {
		errChan <- err
		return
	}

	for _, log := range logs {
		logging.Logger.Infof("get log: %d, %s, %s", log.BlockNumber, log.Topics[0].String(), log.TxHash.String())
		relayTx, err := ParseMocaVCRelayPackage(&l.vcCrossChainAbi,
			&log, uint64(block.Header.Time.Unix()),
			sdk.ChainID(l.config.MocaConfig.ChainId),
			sdk.ChainID(l.config.BSCConfig.ChainId),
		)
		if err != nil {
			logging.Logger.Errorf("failed to parse event log, txHash=%s, err=%s", log.TxHash, err.Error())
			continue
		}
		txChan <- relayTx
	}
}

func (l *MocaListener) monitorEndBlockEvents(height uint64, endBlockEvents []abci.Event, txChan chan *model.MocaRelayTransaction, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, e := range endBlockEvents {
		if e.Type == MocaEventTypeCrossChain {
			relayTx, err := constructRelayTx(e, height)
			if err != nil {
				errChan <- err
				return
			}
			if relayTx.DestChainId != l.destChainId() {
				break
			}
			txChan <- relayTx
		}
	}
}

func (l *MocaListener) monitorValidators(block *tmtypes.Block, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := l.monitorValidatorsHelper(block); err != nil {
		errChan <- err
	}
}

func (l *MocaListener) monitorValidatorsHelper(block *tmtypes.Block) error {
	lightClientLatestHeight, err := l.bscExecutor.GetLightClientLatestHeight()
	if err != nil {
		return err
	}
	nextHeight := uint64(block.Height)
	// happen when re-process block
	if nextHeight <= lightClientLatestHeight {
		return nil
	}

	latestSyncedLightBlockTx, err := l.DaoManager.MocaDao.GetLatestSyncedTransaction()
	if err != nil {
		return err
	}
	latestValidatorsHashFromDB, err := hex.DecodeString(latestSyncedLightBlockTx.ValidatorsHash)
	if err != nil {
		return err
	}

	if bytes.Equal(block.ValidatorsHash[:], latestValidatorsHashFromDB) {
		return nil
	}
	nextValidators, err := l.mocaExecutor.QueryValidatorsAtHeight(nextHeight)
	if err != nil {
		return err
	}

	curValidators, err := l.mocaExecutor.QueryValidatorsAtHeight(nextHeight - 1)
	if err != nil {
		return err
	}

	if len(nextValidators) != len(curValidators) {
		if err = l.sync(nextHeight, block.ValidatorsHash.String()); err != nil {
			return err
		}
		return nil
	}
	for idx, nextVal := range nextValidators {
		curVal := curValidators[idx]

		if !bytes.Equal(nextVal.Address.Bytes(), curVal.Address.Bytes()) ||
			!bytes.Equal(nextVal.BlsKey, curVal.BlsKey) ||
			!bytes.Equal(nextVal.RelayerAddress, curVal.RelayerAddress) {

			if err = l.sync(nextHeight, block.ValidatorsHash.String()); err != nil {
				return err
			}
			break
		}
	}
	return nil
}

func (l *MocaListener) calNextHeight() (uint64, error) {
	latestPolledBlock, err := l.getLatestPolledBlock()
	if err != nil {
		return 0, fmt.Errorf("failed to get latest block from db, error: %s", err.Error())
	}
	latestPolledBlockHeight := latestPolledBlock.Height

	nextHeight := l.config.MocaConfig.StartHeight
	if nextHeight <= latestPolledBlockHeight {
		nextHeight = latestPolledBlockHeight + 1
	}

	latestBlockHeight, err := l.mocaExecutor.GetLatestBlockHeight()
	if err != nil {
		return 0, fmt.Errorf("failed to get latest block height, error: %s", err.Error())
	}
	// pauses relayer for a bit since it already caught the newest block
	if int64(nextHeight) >= int64(latestBlockHeight) {
		time.Sleep(common.ListenerPauseTime)
		return nextHeight, nil
	}
	return nextHeight, nil
}

func (l *MocaListener) sync(nextHeight uint64, validatorsHash string) error {
	logging.Logger.Infof("syncing tendermint light block at height %d", nextHeight)
	txHash, err := l.bscExecutor.SyncTendermintLightBlock(nextHeight)
	if err != nil {
		return fmt.Errorf("failed to sync light block at height=%d, err=%s", nextHeight, err.Error())
	}
	t := &model.SyncLightBlockTransaction{
		ValidatorsHash: validatorsHash,
		Height:         nextHeight,
		TxHash:         txHash.String(),
	}
	if err = l.DaoManager.MocaDao.SaveSyncLightBlockTransaction(t); err != nil {
		return fmt.Errorf("failed to save sync light block transaction to DB, err=%s", err.Error())
	}
	logging.Logger.Infof("synced tendermint light block at height %d with txHash %s", nextHeight, txHash.String())
	time.Sleep(common.SleepTimeAfterSyncLightBlock)
	return nil
}

func constructRelayTx(event abci.Event, height uint64) (*model.MocaRelayTransaction, error) {
	relayTx := model.MocaRelayTransaction{}
	for _, attr := range event.Attributes {
		switch attr.Key {
		case "channel_id":
			chanelId, err := strconv.ParseInt(attr.Value, 10, 8)
			if err != nil {
				return nil, err
			}
			relayTx.ChannelId = uint8(chanelId)
		case "src_chain_id":
			srcChainId, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return nil, err
			}
			relayTx.SrcChainId = uint32(srcChainId)
		case "dest_chain_id":
			destChainId, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return nil, err
			}
			relayTx.DestChainId = uint32(destChainId)
		case "package_load":
			payloadStr, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			relayTx.PayLoad = payloadStr
		case "sequence":
			seq, err := util.QuotedStrToIntWithBitSize(attr.Value, 64)
			if err != nil {
				return nil, err
			}
			relayTx.Sequence = seq
		case "package_type":
			packType, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return nil, err
			}
			relayTx.PackageType = uint32(packType)
		case "timestamp":
			ts, err := util.QuotedStrToIntWithBitSize(attr.Value, 64)
			if err != nil {
				return nil, err
			}
			relayTx.TxTime = int64(ts)
		case "relayer_fee":
			feeStr, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			relayTx.RelayerFee = feeStr
		case "ack_relayer_fee":
			feeStr, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			relayTx.AckRelayerFee = feeStr
		default:
			logging.Logger.Errorf("unexpected attr, key is %s", attr.Key)
		}
	}
	relayTx.Status = db.Saved
	relayTx.Height = height
	relayTx.UpdatedTime = time.Now().Unix()
	return &relayTx, nil
}

func (l *MocaListener) querySBTCrossChainLogs(height uint64) ([]ethtypes.Log, error) {
	client := l.mocaExecutor.GetEthClient()
	topics := [][]ethcommon.Hash{{l.getSBTCrossChainPackageEventHash()}}
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(height)),
		ToBlock:   big.NewInt(int64(height)),
		Topics:    topics,
		Addresses: []ethcommon.Address{l.getSBTCrossChainContractAddress()},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query moca SBT cross chain logs, err=%s", err.Error())
	}
	return logs, nil
}

func (l *MocaListener) queryVCCrossChainLogs(height uint64) ([]ethtypes.Log, error) {
	client := l.mocaExecutor.GetEthClient()
	topics := [][]ethcommon.Hash{{l.getVCCrossChainPackageEventHash()}}
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(height)),
		ToBlock:   big.NewInt(int64(height)),
		Topics:    topics,
		Addresses: []ethcommon.Address{l.getVCCrossChainContractAddress()},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query moca VC cross chain logs, err=%s", err.Error())
	}
	return logs, nil
}

func (l *MocaListener) getSBTCrossChainPackageEventHash() ethcommon.Hash {
	return ethcommon.HexToHash(MocaSBTCrossChainPackageEventHex)
}

func (l *MocaListener) getVCCrossChainPackageEventHash() ethcommon.Hash {
	return ethcommon.HexToHash(VCCrossChainPackageEventHex)
}

func (l *MocaListener) getSBTCrossChainContractAddress() ethcommon.Address {
	return ethcommon.HexToAddress(l.config.RelayConfig.SrcMocaSBTContractAddr)
}

func (l *MocaListener) getVCCrossChainContractAddress() ethcommon.Address {
	return ethcommon.HexToAddress(l.config.RelayConfig.SrcMocaVCContractAddr)
}

func (l *MocaListener) PurgeLoop() {
	ticker := time.NewTicker(PurgeJobInterval)
	for range ticker.C {
		latestGnfdBlock, err := l.DaoManager.MocaDao.GetLatestBlock()
		if err != nil {
			logging.Logger.Errorf("failed to get latest DB BSC block, err=%s", err.Error())
			continue
		}
		threshHold := int64(latestGnfdBlock.Height) - NumOfHistoricalBlocks
		if threshHold <= 0 {
			continue
		}
		if err = l.DaoManager.MocaDao.DeleteBlocksBelowHeight(threshHold); err != nil {
			logging.Logger.Errorf("failed to delete gnfd blocks, err=%s", err.Error())
			continue
		}
		exists, err := l.DaoManager.MocaDao.ExistsUnprocessedTransaction(threshHold)
		if err != nil || exists {
			continue
		}
		if err = l.DaoManager.MocaDao.DeleteTransactionsBelowHeightWithLimit(threshHold, DeletionLimit); err != nil {
			logging.Logger.Errorf("failed to delete gnfd transactions, err=%s", err.Error())
			continue
		}
		var eventType votepool.EventType
		if l.config.BSCConfig.IsOpCrossChain() {
			eventType = votepool.ToOpCrossChainEvent
		} else {
			eventType = votepool.ToBscCrossChainEvent
		}
		if err = l.DaoManager.VoteDao.DeleteVotesBelowHeightWithLimit(threshHold, uint32(eventType), DeletionLimit); err != nil {
			logging.Logger.Errorf("failed to delete votes, err=%s", err.Error())
		}
	}
}

func (l *MocaListener) destChainId() uint32 {
	return uint32(l.config.BSCConfig.ChainId)
}
