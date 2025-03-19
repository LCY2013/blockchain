package repo

import "github.com/LCY2013/blockchain/exchange/ff-coin/jobcenter/internal/model"

type BtcTransactionRepo interface {
	FindByTxId(txId string) (*model.BitCoinTransaction, error)
	Save(bt *model.BitCoinTransaction) error
}
