package dao

import (
	"context"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/msdb"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/msdb/gorms"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market/internal/model"
	"gorm.io/gorm"
)

type CoinDao struct {
	conn *gorms.GormConn
}

func (d *CoinDao) FindById(ctx context.Context, id int64) (*model.Coin, error) {
	session := d.conn.Session(ctx)
	coin := &model.Coin{}
	err := session.Model(&model.Coin{}).Where("id=?", id).Take(coin).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return coin, err
}

func (d *CoinDao) FindAll(ctx context.Context) (list []*model.Coin, err error) {
	session := d.conn.Session(ctx)
	err = session.Model(&model.Coin{}).Find(&list).Error
	return
}

func (d *CoinDao) FindByUnit(ctx context.Context, unit string) (*model.Coin, error) {
	session := d.conn.Session(ctx)
	coin := &model.Coin{}
	err := session.Model(&model.Coin{}).Where("unit=?", unit).Take(coin).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return coin, err
}

func NewCoinDao(db *msdb.MsDB) *CoinDao {
	return &CoinDao{
		conn: gorms.New(db.Conn),
	}
}
