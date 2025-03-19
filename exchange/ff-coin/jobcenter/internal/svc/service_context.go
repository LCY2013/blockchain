package svc

import (
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/ucenter/ucclient"
	"github.com/LCY2013/blockchain/exchange/ff-coin/jobcenter/internal/config"
	"github.com/LCY2013/blockchain/exchange/ff-coin/jobcenter/internal/database"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	MongoClient    *database.MongoClient
	KafkaClient    *database.KafkaClient
	Cache          cache.Cache
	AssetRpc       ucclient.Asset
	BitCoinAddress string
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := database.NewKafkaClient(c.Kafka)
	client.StartWrite()
	redisCache := cache.New(
		c.CacheRedis,
		nil,
		cache.NewStat("ffcoin"),
		nil,
		func(o *cache.Options) {})
	return &ServiceContext{
		Config:         c,
		MongoClient:    database.ConnectMongo(c.Mongo),
		KafkaClient:    client,
		Cache:          redisCache,
		AssetRpc:       ucclient.NewAsset(zrpc.MustNewClient(c.UCenterRpc)),
		BitCoinAddress: c.Bitcoin.Address,
	}
}
