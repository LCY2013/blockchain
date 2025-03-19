package domain

import (
	"context"
	"errors"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/msdb"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market/internal/dao"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market/internal/model"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market/internal/repo"
)

type CoinDomain struct {
	coinRepo repo.CoinRepo
}

func (d *CoinDomain) FindCoinInfo(ctx context.Context, unit string) (*model.Coin, error) {
	coin, err := d.coinRepo.FindByUnit(ctx, unit)
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, errors.New("not support this coin:" + unit)
	}
	return coin, nil
}
func (d *CoinDomain) FindCoinId(ctx context.Context, id int64) (*model.Coin, error) {
	coin, err := d.coinRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, errors.New("not support this coin")
	}
	return coin, nil
}

func (d *CoinDomain) FindAll(ctx context.Context) ([]*model.Coin, error) {
	return d.coinRepo.FindAll(ctx)
}

func NewCoinDomain(db *msdb.MsDB) *CoinDomain {
	return &CoinDomain{
		coinRepo: dao.NewCoinDao(db),
	}
}
