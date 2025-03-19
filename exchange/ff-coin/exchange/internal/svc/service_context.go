package svc

import (
	"github.com/LCY2013/blockchain/exchange/ff-coin/exchange/internal/config"
	"github.com/LCY2013/blockchain/exchange/ff-coin/exchange/internal/consumer"
	"github.com/LCY2013/blockchain/exchange/ff-coin/exchange/internal/database"
	"github.com/LCY2013/blockchain/exchange/ff-coin/exchange/internal/processor"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/msdb"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/market/mclient"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/ucenter/ucclient"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Cache       cache.Cache
	Db          *msdb.MsDB
	MongoClient *database.MongoClient
	MemberRpc   ucclient.Member
	MarketRpc   mclient.Market
	AssetRpc    ucclient.Asset
	KafkaClient *database.KafkaClient
}

func (sc *ServiceContext) init() {
	factory := processor.NewCoinTradeFactory()
	factory.Init(sc.MarketRpc, sc.KafkaClient, sc.Db)
	kafkaConsumer := consumer.NewKafkaConsumer(sc.KafkaClient, factory, sc.Db)
	kafkaConsumer.Run()
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := cache.New(
		c.CacheRedis,
		nil,
		cache.NewStat("market"),
		nil,
		func(o *cache.Options) {})
	kafkaClient := database.NewKafkaClient(c.Kafka)
	client, _ := zrpc.NewClient(c.UCenterRpc)
	newClient, _ := zrpc.NewClient(c.MarketRpc)
	s := &ServiceContext{
		Config:      c,
		Cache:       redisCache,
		Db:          database.ConnMysql(c.Mysql),
		MongoClient: database.ConnectMongo(c.Mongo),
		MemberRpc:   ucclient.NewMember(client),
		MarketRpc:   mclient.NewMarket(newClient),
		AssetRpc:    ucclient.NewAsset(client),
		KafkaClient: kafkaClient,
	}
	s.init()
	return s
}
