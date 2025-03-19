package config

import (
	"github.com/LCY2013/blockchain/exchange/ff-coin/jobcenter/internal/database"
	"github.com/LCY2013/blockchain/exchange/ff-coin/jobcenter/internal/logic"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Okx        logic.OkxConfig
	Mongo      database.MongoConfig
	Kafka      database.KafkaConfig
	CacheRedis cache.CacheConf
	UCenterRpc zrpc.RpcClientConf
	Bitcoin    logic.BitCoinConfig
}
