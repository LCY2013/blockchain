package domain

import (
	"context"
	"github.com/LCY2013/blockchain/exchange/ff-coin/jobcenter/internal/dao"
	"github.com/LCY2013/blockchain/exchange/ff-coin/jobcenter/internal/database"
	"github.com/LCY2013/blockchain/exchange/ff-coin/jobcenter/internal/model"
	"github.com/LCY2013/blockchain/exchange/ff-coin/jobcenter/internal/repo"
	"log"
)

type KlineDomain struct {
	klineRepo repo.KlineRepo
}

func (d *KlineDomain) SaveBatch(data [][]string, symbol string, period string) {
	klines := make([]*model.Kline, len(data))
	for i, v := range data {
		klines[i] = model.NewKline(v, period)
	}
	err := d.klineRepo.DeleteGtTime(context.Background(), klines[len(data)-1].Time, symbol, period)
	if err != nil {
		log.Println(err)
		return
	}
	err = d.klineRepo.SaveBatch(context.Background(), klines, symbol, period)
	if err != nil {
		log.Println(err)
	}
}

func NewKlineDomain(cli *database.MongoClient) *KlineDomain {
	return &KlineDomain{
		klineRepo: dao.NewKlineDao(cli.Db),
	}
}
