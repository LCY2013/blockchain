package main

import (
	"flag"
	"fmt"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/ucenter/types/asset"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/ucenter/types/login"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/ucenter/types/member"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/ucenter/types/register"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/ucenter/types/withdraw"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter/internal/config"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter/internal/server"
	"github.com/LCY2013/blockchain/exchange/ff-coin/ucenter/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

func main() {
	flag.Parse()
	//日志的打印格式替换一下
	logx.MustSetup(logx.LogConf{Stat: false, Encoding: "plain"})
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		register.RegisterRegisterServer(grpcServer, server.NewRegisterServer(ctx))
		login.RegisterLoginServer(grpcServer, server.NewLoginServer(ctx))
		asset.RegisterAssetServer(grpcServer, server.NewAssetServer(ctx))
		member.RegisterMemberServer(grpcServer, server.NewMemberServer(ctx))
		withdraw.RegisterWithdrawServer(grpcServer, server.NewWithdrawServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
