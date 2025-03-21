// Code generated by goctl. DO NOT EDIT.
// Source: register.proto

package server

import (
	"context"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/market/types/rate"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market/internal/logic"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market/internal/svc"
)

type ExchangeRateServer struct {
	svcCtx *svc.ServiceContext
	rate.UnimplementedExchangeRateServer
}

func (e *ExchangeRateServer) UsdRate(ctx context.Context, req *rate.RateReq) (*rate.RateRes, error) {
	l := logic.NewExchangeRateLogic(ctx,e.svcCtx)
	return l.UsdRate(req)
}

func NewExchangeRateServer(svcCtx *svc.ServiceContext) *ExchangeRateServer {
	return &ExchangeRateServer{
		svcCtx: svcCtx,
	}
}
