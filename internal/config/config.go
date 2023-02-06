package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConf cache.CacheConf
	MongoConf struct {
		Source      string
		Database    string
		CollComment string
		CollHistory string
	}
	MqConf struct {
		NameServer []string
		Retry      int
		GroupName  string
	}
}
