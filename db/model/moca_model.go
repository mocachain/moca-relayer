package model

import (
	"gorm.io/gorm"

	"github.com/mocachain/moca-relayer/db"
)

type MocaBlock struct {
	Id        int64
	Chain     string
	Height    uint64 `gorm:"NOT NULL;index:idx_moca_block_height"`
	BlockTime int64  `gorm:"NOT NULL"`
}

func (*MocaBlock) TableName() string {
	return "moca_block"
}

type MocaRelayTransaction struct {
	Id            int64
	SrcChainId    uint32 `gorm:"NOT NULL"`
	DestChainId   uint32 `gorm:"NOT NULL"`
	ChannelId     uint8  `gorm:"NOT NULL;index:idx_moca_relay_transaction_channel_seq_status"`
	Sequence      uint64 `gorm:"NOT NULL;index:idx_moca_relay_transaction_channel_seq_status"`
	PackageType   uint32 `gorm:"NOT NULL"`
	Height        uint64 `gorm:"NOT NULL;index:idx_moca_relay_transaction_height_status"`
	PayLoad       string `gorm:"type:text"`
	RelayerFee    string `gorm:"NOT NULL"`
	AckRelayerFee string `gorm:"NOT NULL"`
	TxHash        string
	ClaimedTxHash string
	Status        db.TxStatus `gorm:"NOT NULL;index:idx_moca_relay_transaction_channel_seq_status;idx_moca_relay_transaction_height_status"`
	TxTime        int64       `gorm:"NOT NULL"`
	UpdatedTime   int64       `gorm:"NOT NULL"`
}

func (*MocaRelayTransaction) TableName() string {
	return "moca_relay_transaction"
}

type SyncLightBlockTransaction struct {
	Id             int64
	ValidatorsHash string `gorm:"NOT NULL"`
	Height         uint64 `gorm:"NOT NULL;index:idx_sync_light_block_transaction_height"`
	TxHash         string `gorm:"NOT NULL"`
}

func (*SyncLightBlockTransaction) TableName() string {
	return "sync_light_block_transaction"
}

func InitMocaTables(db *gorm.DB) {
	if !db.Migrator().HasTable(&MocaBlock{}) {
		err := db.Migrator().CreateTable(&MocaBlock{})
		if err != nil {
			panic(err)
		}
	}

	if !db.Migrator().HasTable(&MocaRelayTransaction{}) {
		err := db.Migrator().CreateTable(&MocaRelayTransaction{})
		if err != nil {
			panic(err)
		}
	}

	if !db.Migrator().HasTable(&SyncLightBlockTransaction{}) {
		err := db.Migrator().CreateTable(&SyncLightBlockTransaction{})
		if err != nil {
			panic(err)
		}
	}
}
