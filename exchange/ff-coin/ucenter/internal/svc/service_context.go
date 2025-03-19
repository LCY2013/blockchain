package svc

import (
	"github.com/LCY2013/blockchain/exchange/ff-coin/ffcoin-common/msdb"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/exchange/eclient"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/market/mclient"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter/internal/config"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter/internal/consumer"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter/internal/database"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	Cache          cache.Cache
	Db             *msdb.MsDB
	MarketRpc      mclient.Market
	KafkaCli       *database.KafkaClient
	BitcoinAddress string
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := cache.New(
		c.CacheRedis,
		nil,
		cache.NewStat("ffcoin"),
		nil,
		func(o *cache.Options) {})
	mysql := database.ConnMysql(c.Mysql.DataSource)
	cli := database.NewKafkaClient(c.Kafka)
	cli.StartRead("add-exchange-order")
	order := eclient.NewOrder(zrpc.MustNewClient(c.ExchangeRpc))
	conf := c.CacheRedis[0].RedisConf
	newRedis := redis.MustNewRedis(conf)
	go consumer.ExchangeOrderAdd(newRedis, cli, order, mysql)
	completeCli := cli.StartReadNew("exchange_order_complete_update_success")
	go consumer.ExchangeOrderComplete(newRedis, completeCli, mysql)
	btCli := cli.StartReadNew("BTC_TRANSACTION")
	go consumer.BitCoinTransaction(newRedis, btCli, mysql)
	withdrawCli := cli.StartReadNew("withdraw")
	go consumer.WithdrawConsumer(withdrawCli, mysql, c.Bitcoin.Address)
	return &ServiceContext{
		Config:         c,
		Cache:          redisCache,
		Db:             mysql,
		MarketRpc:      mclient.NewMarket(zrpc.MustNewClient(c.MarketRpc)),
		KafkaCli:       cli,
		BitcoinAddress: c.Bitcoin.Address,
	}
}
