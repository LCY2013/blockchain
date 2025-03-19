package logic

import (
	"context"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/market/types/rate"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market-api/internal/svc"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market-api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type ExchangeRateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (l *ExchangeRateLogic) UsdRate(req *types.RateRequest) (*types.RateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rateRes, err := l.svcCtx.ExchangeRateRpc.UsdRate(ctx, &rate.RateReq{
		Unit: req.Unit,
		Ip:   req.Ip,
	})
	if err != nil {
		return nil, err
	}
	return &types.RateResponse{
		Rate: rateRes.Rate,
	}, nil
}

func NewExchangeRateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeRateLogic {
	return &ExchangeRateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
