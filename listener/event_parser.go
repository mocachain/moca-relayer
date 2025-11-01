package listener

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/mocachain/moca-relayer/common"
	"github.com/mocachain/moca-relayer/contract/universalVerifier"
	"github.com/mocachain/moca-relayer/contract/mocacrosschainupgradeable"
	"github.com/mocachain/moca-relayer/db"
	"github.com/mocachain/moca-relayer/db/model"
	rtypes "github.com/mocachain/moca-relayer/types"
)

func ParseRelayPackage(abi *abi.ABI, log *types.Log, timestamp uint64, mocaChainId, bscChainId sdk.ChainID) (*model.BscRelayPackage, error) {
	ev, err := parseCrossChainPackageEvent(abi, log)
	if err != nil {
		return nil, err
	}
	if sdk.ChainID(ev.SrcChainId) != bscChainId || sdk.ChainID(ev.DstChainId) != mocaChainId {
		return nil, fmt.Errorf("event log's chain id(s) not expected, SrcChainId=%d, DstChainId=%d", ev.SrcChainId, ev.DstChainId)
	}
	var p model.BscRelayPackage
	p.OracleSequence = ev.OracleSequence
	p.PackageSequence = ev.PackageSequence
	p.ChannelId = ev.ChannelId
	p.TxHash = log.TxHash.String()
	p.TxIndex = log.TxIndex
	p.TxTime = int64(timestamp)
	p.UpdatedTime = int64(timestamp)
	p.Height = log.BlockNumber
	p.Status = db.Saved
	p.PayLoad = hex.EncodeToString(ev.Payload)
	return &p, nil
}

func parseCrossChainPackageEvent(abi *abi.ABI, log *types.Log) (*rtypes.CrossChainPackageEvent, error) {
	var ev rtypes.CrossChainPackageEvent

	err := abi.UnpackIntoInterface(&ev, BSCCrossChainPackageEventName, log.Data)
	if err != nil {
		return nil, err
	}
	ev.OracleSequence = big.NewInt(0).SetBytes(log.Topics[1].Bytes()).Uint64()
	ev.PackageSequence = big.NewInt(0).SetBytes(log.Topics[2].Bytes()).Uint64()
	ev.ChannelId = uint8(big.NewInt(0).SetBytes(log.Topics[3].Bytes()).Uint64())
	return &ev, nil
}

func ParseMocaSBTRelayPackage(abi *abi.ABI, log *types.Log, timestamp uint64, mocaChainId, bscChainId sdk.ChainID) (*model.MocaRelayTransaction, error) {
	ev, err := parseMocaSBTCrossChainPackageEvent(abi, log)
	if err != nil {
		return nil, err
	}
	if sdk.ChainID(ev.SrcChainId) != mocaChainId || sdk.ChainID(ev.DestChainId) != bscChainId {
		return nil, fmt.Errorf("mocasbt event log's chain id(s) not expected, SrcChainId=%d, DstChainId=%d, ChannelId=%d, Sequence=%d",
			ev.SrcChainId, ev.DestChainId, ev.ChannelId, ev.Sequence.Int64())
	}
	var p model.MocaRelayTransaction

	p.ChannelId = uint8(ev.ChannelId)
	p.SrcChainId = ev.SrcChainId
	p.DestChainId = ev.DestChainId
	p.Sequence = ev.Sequence.Uint64()
	p.PackageType = uint32(sdk.SynCrossChainPackageType)
	p.TxHash = log.TxHash.String()
	p.TxTime = int64(timestamp)
	// Assign RelayerFee based on the chainid of blockchain.
	// if ev.DestChainId == "" {}
	p.RelayerFee = common.DefaultBscMirrorMocaSBTRelayerFee
	p.AckRelayerFee = common.DefaultBscMirrorMocaSBTAckRelayerFee
	p.UpdatedTime = time.Now().Unix()
	p.Height = log.BlockNumber
	p.Status = db.Saved
	p.PayLoad = hex.EncodeToString(ev.Payload)
	return &p, nil
}

func parseMocaSBTCrossChainPackageEvent(abi *abi.ABI, log *types.Log) (*mocacrosschainupgradeable.KYCDataLibEventData, error) {
	var ev mocacrosschainupgradeable.KYCDataLibEventData

	err := abi.UnpackIntoInterface(&ev, MocaSBTCrossChainPackageEventName, log.Data)
	if err != nil {
		return nil, err
	}
	ev.DestChainId = uint32(big.NewInt(0).SetBytes(log.Topics[1].Bytes()).Uint64())
	ev.ChannelId = uint32(big.NewInt(0).SetBytes(log.Topics[2].Bytes()).Uint64())
	ev.Sequence = big.NewInt(0).SetBytes(log.Topics[3].Bytes())
	return &ev, nil
}

func ParseMocaVCRelayPackage(abi *abi.ABI, log *types.Log, timestamp uint64, mocaChainId, bscChainId sdk.ChainID) (*model.MocaRelayTransaction, error) {
	ev, err := parseMocaVCCrossChainPackageEvent(abi, log)
	if err != nil {
		return nil, err
	}
	if sdk.ChainID(ev.SrcChainId) != mocaChainId || sdk.ChainID(ev.DestChainId) != bscChainId {
		return nil, fmt.Errorf("mocavc event log's chain id(s) not expected, SrcChainId=%d, DstChainId=%d, ChannelId=%d, Sequence=%d",
			ev.SrcChainId, ev.DestChainId, ev.ChannelId, ev.Sequence.Int64())
	}
	var p model.MocaRelayTransaction

	p.ChannelId = uint8(ev.ChannelId)
	p.SrcChainId = ev.SrcChainId
	p.DestChainId = ev.DestChainId
	p.Sequence = ev.Sequence.Uint64()
	p.PackageType = uint32(sdk.SynCrossChainPackageType)
	p.TxHash = log.TxHash.String()
	p.TxTime = int64(timestamp)
	// Assign RelayerFee based on the chainid of blockchain.
	// if ev.DestChainId == "" {}
	p.RelayerFee = common.DefaultBscMirrorMocaSBTRelayerFee
	p.AckRelayerFee = common.DefaultBscMirrorMocaSBTAckRelayerFee
	p.UpdatedTime = time.Now().Unix()
	p.Height = log.BlockNumber
	p.Status = db.Saved
	p.PayLoad = hex.EncodeToString(ev.CrossDataBytes)
	return &p, nil
}

func parseMocaVCCrossChainPackageEvent(abi *abi.ABI, log *types.Log) (*universalVerifier.VCDataLibEventData, error) {
	var ev universalVerifier.VCDataLibEventData

	err := abi.UnpackIntoInterface(&ev, MocaVCCrossChainPackageEventName, log.Data)
	if err != nil {
		return nil, err
	}
	ev.DestChainId = uint32(big.NewInt(0).SetBytes(log.Topics[1].Bytes()).Uint64())
	ev.ChannelId = uint32(big.NewInt(0).SetBytes(log.Topics[2].Bytes()).Uint64())
	ev.Sequence = big.NewInt(0).SetBytes(log.Topics[3].Bytes())
	return &ev, nil
}
