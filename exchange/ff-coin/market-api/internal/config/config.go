package config

import (
	"github.com/LCY2013/blockchain/exchange/ff-coin/market-api/internal/database"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Prefix    string
	MarketRpc zrpc.RpcClientConf
	Kafka     database.KafkaConfig
}
