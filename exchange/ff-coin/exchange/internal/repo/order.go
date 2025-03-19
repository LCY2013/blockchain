package repo

import (
	"context"
	"github.com/LCY2013/blockchain/exchange/ff-coin/exchange/internal/model"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/msdb"
)

type ExchangeOrderRepo interface {
	FindOrderHistory(ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrder, total int64, err error)
	FindOrderCurrent(ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrder, total int64, err error)
	FindCurrentTradingCount(ctx context.Context, id int64, symbol string, direction int) (int64, error)
	Save(ctx context.Context, conn msdb.DbConn, order *model.ExchangeOrder) error
	FindOrderByOrderId(ctx context.Context, orderId string) (*model.ExchangeOrder, error)
	UpdateStatusCancel(ctx context.Context, orderId string) error
	UpdateOrderStatusTrading(ctx context.Context, orderId string) error
	FindOrderListBySymbol(ctx context.Context, symbol string, status int) ([]*model.ExchangeOrder, error)
	UpdateOrderComplete(ctx context.Context, orderId string, tradedAmount float64, turnover float64, status int) error
}
