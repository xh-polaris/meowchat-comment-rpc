package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Cache cache.CacheConf
	Mongo struct {
		URL string
		DB  string
	}
	MqConf struct {
		NameServer []string
		Retry      int
		GroupName  string
	}
}
