package domain

import (
	"github.com/LCY2013/blockchain/exchange/ff-coin/jobcenter/internal/dao"
	"github.com/LCY2013/blockchain/exchange/ff-coin/jobcenter/internal/database"
	"github.com/LCY2013/blockchain/exchange/ff-coin/jobcenter/internal/model"
	"github.com/LCY2013/blockchain/exchange/ff-coin/jobcenter/internal/repo"
)

type BitCoinDomain struct {
	btcTransactionRepo repo.BtcTransactionRepo
}

func (d *BitCoinDomain) Recharge(
	txId string,
	value float64,
	address string,
	time int64,
	blockhash string) error {
	bitCoinTransaction, err := d.btcTransactionRepo.FindByTxId(txId)
	if err != nil {
		return err
	}
	if bitCoinTransaction == nil {
		bt := &model.BitCoinTransaction{}
		bt.Type = model.RECHARGE
		bt.Time = time
		bt.BlockHash = blockhash
		bt.Value = value
		bt.TxId = txId
		bt.Address = address
		err := d.btcTransactionRepo.Save(bt)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBitCoinDomain(cli *database.MongoClient) *BitCoinDomain {
	return &BitCoinDomain{
		btcTransactionRepo: dao.NewBtcTransactionDao(cli.Db),
	}
}
