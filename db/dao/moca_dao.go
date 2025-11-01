package dao

import (
	"database/sql"
	"time"

	"gorm.io/gorm"

	"github.com/mocachain/moca-relayer/db"
	"github.com/mocachain/moca-relayer/db/model"
	"github.com/mocachain/moca-relayer/types"
)

type MocaDao struct {
	DB *gorm.DB
}

func NewMocaDao(db *gorm.DB) *MocaDao {
	return &MocaDao{
		DB: db,
	}
}

func (d *MocaDao) GetLatestBlock() (*model.MocaBlock, error) {
	block := model.MocaBlock{}
	err := d.DB.Model(model.MocaBlock{}).Order("height desc").Take(&block).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &block, nil
}

func (d *MocaDao) GetTransactionsByStatusWithLimit(s db.TxStatus, limit int64) ([]*model.MocaRelayTransaction, error) {
	txs := make([]*model.MocaRelayTransaction, 0)
	err := d.DB.Where("status = ? ", s).Order("height asc").Limit(int(limit)).Find(&txs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return txs, nil
}

func (d *MocaDao) GetLeastSavedTransactionHeight() (uint64, error) {
	var result sql.NullInt64
	res := d.DB.Table("moca_relay_transaction").Select("MIN(height)").Where("status = ?", db.Saved)
	err := res.Row().Scan(&result)
	if err != nil {
		return 0, err
	}
	return uint64(result.Int64), nil
}

func (d *MocaDao) GetTransactionByChannelIdAndSequence(channelId types.ChannelId, sequence uint64) (*model.MocaRelayTransaction, error) {
	tx := model.MocaRelayTransaction{}
	err := d.DB.Where("channel_id = ? and sequence = ?", channelId, sequence).Find(&tx).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tx, nil
}

// for mocaSBT
// func (d *MocaDao) GetLatestSequenceByChannelId(channelId types.ChannelId) (int64, error) {
// 	var result sql.NullInt64
// 	res := d.DB.Table("moca_relay_transaction").Select("MAX(sequence)").Where("channel_id = ?", channelId)
// 	err := res.Row().Scan(&result)
// 	if err != nil {
// 		return 0, err
// 	}
// 	if !result.Valid {
// 		return -1, nil
// 	}
// 	return result.Int64 + 1, nil
// }

func (d *MocaDao) GetLatestSequenceByChannelIdAndStatus(channelId types.ChannelId, status db.TxStatus) (int64, error) {
	var result sql.NullInt64
	res := d.DB.Table("moca_relay_transaction").Select("MAX(sequence)").Where("channel_id = ? and status = ?", channelId, status)
	err := res.Row().Scan(&result)
	if err != nil {
		return 0, err
	}
	if !result.Valid {
		return -1, nil
	}
	return result.Int64 + 1, nil
}

func (d *MocaDao) UpdateTransactionStatus(id int64, status db.TxStatus) error {
	err := d.DB.Model(model.MocaRelayTransaction{}).Where("id = ?", id).Updates(
		model.MocaRelayTransaction{Status: status, UpdatedTime: time.Now().Unix()}).Error
	return err
}

func UpdateTransactionStatus(dbTx *gorm.DB, id int64, status db.TxStatus) error {
	err := dbTx.Model(model.MocaRelayTransaction{}).Where("id = ?", id).Updates(
		model.MocaRelayTransaction{Status: status, UpdatedTime: time.Now().Unix()}).Error
	return err
}

func (d *MocaDao) UpdateTransactionClaimedTxHash(id int64, claimedTxHash string) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(model.MocaRelayTransaction{}).Where("id = ?", id).Updates(
			model.MocaRelayTransaction{UpdatedTime: time.Now().Unix(), ClaimedTxHash: claimedTxHash}).Error
	})
}

func (d *MocaDao) UpdateTransactionStatusAndClaimedTxHash(id int64, status db.TxStatus, claimedTxHash string) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(model.MocaRelayTransaction{}).Where("id = ?", id).Updates(
			model.MocaRelayTransaction{Status: status, UpdatedTime: time.Now().Unix(), ClaimedTxHash: claimedTxHash}).Error
	})
}

func (d *MocaDao) UpdateBatchTransactionStatusToDelivered(seq uint64) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(model.MocaRelayTransaction{}).Where("sequence < ? and status = 2", seq).Updates(
			model.MocaRelayTransaction{Status: db.Delivered, UpdatedTime: time.Now().Unix()}).Error
	})
}

func (d *MocaDao) SaveBlockAndBatchTransactions(b *model.MocaBlock, txs []*model.MocaRelayTransaction) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		err := dbTx.Create(b).Error
		if err != nil {
			return err
		}

		if len(txs) != 0 {
			err := dbTx.Create(txs).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (d *MocaDao) SaveSyncLightBlockTransaction(t *model.SyncLightBlockTransaction) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Create(t).Error
	})
}

func (d *MocaDao) GetLatestSyncedTransaction() (*model.SyncLightBlockTransaction, error) {
	tx := model.SyncLightBlockTransaction{}
	err := d.DB.Model(model.SyncLightBlockTransaction{}).Order("height desc").Take(&tx).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tx, nil
}

func (d *MocaDao) DeleteBlocksBelowHeight(threshHold int64) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		err := dbTx.Where("height < ?", threshHold).Delete(model.MocaBlock{}).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (d *MocaDao) DeleteTransactionsBelowHeightWithLimit(threshHold int64, limit int) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		err := dbTx.Where("height < ?", threshHold).Delete(model.MocaRelayTransaction{}).Limit(limit).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (d *MocaDao) ExistsUnprocessedTransaction(threshHold int64) (bool, error) {
	tx := model.MocaRelayTransaction{}
	err := d.DB.Model(model.MocaRelayTransaction{}).Where("status = ? or status = ? and height < ?", db.Saved, db.SelfVoted, threshHold).Take(&tx).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
