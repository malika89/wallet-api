package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	Rpc zrpc.RpcServerConf
	rest.RestConf
	Consul consul.Conf
	Store  StoreConf
}

type StoreConf struct {
	MysqlDBUri string
	RedisUri   string
}
