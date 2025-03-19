package repo

import (
	"context"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter/internal/model"
)

type MemberAddressRepo interface {
	FindByMemIdAndCoinId(ctx context.Context, memId int64, coinId int64) ([]*model.MemberAddress, error)
}
