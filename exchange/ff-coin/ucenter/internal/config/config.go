package config

import (
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter/internal/database"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql       MysqlConfig
	CacheRedis  cache.CacheConf
	Captcha     CaptchaConf
	JWT         AuthConfig
	MarketRpc   zrpc.RpcClientConf
	ExchangeRpc zrpc.RpcClientConf
	Kafka       database.KafkaConfig
	Bitcoin     BitCoinConfig
}
type BitCoinConfig struct {
	Address string
}
type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}
type CaptchaConf struct {
	Vid string
	Key string
}
type MysqlConfig struct {
	DataSource string
}
