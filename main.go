package main

import (
	"flag"
	"fmt"
	"github.com/malika89/wallet-api/proto"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/malika89/wallet-api/internal/config"
	"github.com/malika89/wallet-api/internal/server"
	"github.com/malika89/wallet-api/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/wallet.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	fmt.Printf("\n====>consul%v \n", c.Consul)
	ctx := svc.NewServiceContext(c)
	grpcServer := zrpc.MustNewServer(c.Rpc, func(grpcServer *grpc.Server) {
		proto.RegisterWalletServer(grpcServer, server.NewWalletServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer grpcServer.Stop()
	_ = consul.RegisterService(c.Rpc.ListenOn, c.Consul)
	fmt.Printf("Starting rpc server at %s...\n", c.Rpc.ListenOn)
	grpcServer.Start()

	httpServer := rest.MustNewServer(c.RestConf)
	defer httpServer.Stop()
	server.RegisterHandlers(httpServer, ctx)
	httpServer.Start()
}
