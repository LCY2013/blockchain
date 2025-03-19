package svc

import (
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/market/mclient"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/ucenter/ucclient"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter-api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UCRegisterRpc ucclient.Register
	UCLoginRpc    ucclient.Login
	UCAssetRpc    ucclient.Asset
	UCMemberRpc   ucclient.Member
	UCWithdrawRpc ucclient.Withdraw
	MarketRpc     mclient.Market
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UCRegisterRpc: ucclient.NewRegister(zrpc.MustNewClient(c.UCenterRpc)),
		UCLoginRpc:    ucclient.NewLogin(zrpc.MustNewClient(c.UCenterRpc)),
		UCAssetRpc:    ucclient.NewAsset(zrpc.MustNewClient(c.UCenterRpc)),
		UCMemberRpc:   ucclient.NewMember(zrpc.MustNewClient(c.UCenterRpc)),
		UCWithdrawRpc: ucclient.NewWithdraw(zrpc.MustNewClient(c.UCenterRpc)),
		MarketRpc:     mclient.NewMarket(zrpc.MustNewClient(c.MarketRpc)),
	}
}
