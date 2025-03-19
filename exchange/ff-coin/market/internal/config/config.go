package config

import (
	"github.com/LCY2013/blockchain/exchange/ff-coin/market/internal/database"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql      database.MysqlConfig
	CacheRedis cache.CacheConf
	Mongo      database.MongoConfig
}
