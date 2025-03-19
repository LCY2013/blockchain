package config

import (
	"github.com/LCY2013/blockchain/exchange/ff-coin/exchange-api/internal/database"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Prefix      string
	ExchangeRpc zrpc.RpcClientConf
	Kafka       database.KafkaConfig
	JWT         AuthConfig
}
type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}
