package repo

import (
	"context"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market/internal/model"
)

type KlineRepo interface {
	FindBySymbol(ctx context.Context, symbol, period string, count int64) ([]*model.Kline, error)
	FindBySymbolTime(ctx context.Context, symbol, period string, from, end int64, s string) ([]*model.Kline, error)
}
