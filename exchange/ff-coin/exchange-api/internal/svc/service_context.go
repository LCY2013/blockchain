package svc

import (
	"github.com/LCY2013/blockchain/exchange/ff-coin/exchange-api/internal/config"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/exchange/eclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	OrderRpc eclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	order := eclient.NewOrder(zrpc.MustNewClient(c.ExchangeRpc))
	return &ServiceContext{
		Config:   c,
		OrderRpc: order,
	}
}
