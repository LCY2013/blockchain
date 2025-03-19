package svc

import (
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/msdb"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market/internal/config"
	"github.com/LCY2013/blockchain/exchange/ff-coin/market/internal/database"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type ServiceContext struct {
	Config      config.Config
	Cache       cache.Cache
	Db          *msdb.MsDB
	MongoClient *database.MongoClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := cache.New(
		c.CacheRedis,
		nil,
		cache.NewStat("market"),
		nil,
		func(o *cache.Options) {})
	return &ServiceContext{
		Config:      c,
		Cache:       redisCache,
		Db:          database.ConnMysql(c.Mysql),
		MongoClient: database.ConnectMongo(c.Mongo),
	}
}
