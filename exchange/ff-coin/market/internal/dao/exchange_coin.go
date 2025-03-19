package dao

import (
	"context"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/msdb"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/msdb/gorms"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market/internal/model"
	"gorm.io/gorm"
)

type ExchangeCoinDao struct {
	conn *gorms.GormConn
}

func (e *ExchangeCoinDao) FindBySymbol(ctx context.Context, symbol string) (*model.ExchangeCoin, error) {
	session := e.conn.Session(ctx)
	data := &model.ExchangeCoin{}
	err := session.Model(&model.ExchangeCoin{}).Where("symbol=?", symbol).Take(data).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return data, err
}

func (e *ExchangeCoinDao) FindVisible(ctx context.Context) (list []*model.ExchangeCoin, err error) {
	session := e.conn.Session(ctx)
	err = session.Model(&model.ExchangeCoin{}).Where("visible=?", 1).Find(&list).Error
	return
}

func NewExchangeCoinDao(db *msdb.MsDB) *ExchangeCoinDao {
	return &ExchangeCoinDao{
		conn: gorms.New(db.Conn),
	}
}
