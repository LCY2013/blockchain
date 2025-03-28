package config

import (
	"github.com/LCY2013/blockchain/exchange/ff-coin/exchange/internal/database"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql      database.MysqlConfig
	CacheRedis cache.CacheConf
	Mongo      database.MongoConfig
	UCenterRpc zrpc.RpcClientConf
	MarketRpc  zrpc.RpcClientConf
	Kafka      database.KafkaConfig
}
