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
		DataBase    string
		CollComment string
		CollHistory string
	}
}
