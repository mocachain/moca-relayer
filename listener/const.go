package listener

import (
	"time"
)

const (
	NumOfHistoricalBlocks             = int64(50000) // NumOfHistoricalBlocks is the number of blocks will be kept in DB, all transactions and votes also kept within this range
	PurgeJobInterval                  = time.Minute * 1
	DeletionLimit                     = 10000
	MocaEventTypeCrossChain           = "cosmos.crosschain.v1.EventCrossChain"
	BSCCrossChainPackageEventName     = "CrossChainPackage"
	MocaSBTCrossChainPackageEventName = "MocaSBTCrossChainPackage"
	MocaVCCrossChainPackageEventName  = "MocaVCCrossChainPackage"
	CrossChainPackageEventHex         = "0x64998dc5a229e7324e622192f111c691edccc3534bbea4b2bd90fbaec936845a"
	MocaSBTCrossChainPackageEventHex  = "0xeae7aa948aa4486965776005e20135ce32c4e9a9bd3704ec53d108056bcba038"
	VCCrossChainPackageEventHex       = "0x2f4796132af44fdc7d0bfd7fe25cd97e2b6e8981ca28b81dfd6e22f065a01c50"
)
