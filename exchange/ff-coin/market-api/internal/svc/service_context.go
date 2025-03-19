package svc

import (
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/market/mclient"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market-api/internal/config"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market-api/internal/database"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market-api/internal/processor"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market-api/internal/ws"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	ExchangeRateRpc mclient.ExchangeRate
	MarketRpc       mclient.Market
	Processor       processor.Processor
}

func NewServiceContext(c config.Config, server *ws.WebsocketServer) *ServiceContext {
	//初始化processor
	kafaCli := database.NewKafkaClient(c.Kafka)
	market := mclient.NewMarket(zrpc.MustNewClient(c.MarketRpc))
	defaultProcessor := processor.NewDefaultProcessor(kafaCli)
	defaultProcessor.Init(market)
	defaultProcessor.AddHandler(processor.NewWebsocketHandler(server))
	return &ServiceContext{
		Config:          c,
		ExchangeRateRpc: mclient.NewExchangeRate(zrpc.MustNewClient(c.MarketRpc)),
		MarketRpc:       market,
		Processor:       defaultProcessor,
	}
}
