package svc

import (
	"github.com/xh-polaris/meowchat-post-rpc/internal/config"
	"github.com/xh-polaris/meowchat-post-rpc/internal/model/mongo/commentcached"
	"github.com/xh-polaris/meowchat-post-rpc/internal/model/mongo/commenthistory"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel commentcached.CommentModel
	HistoryModel commenthistory.HistoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		CommentModel: commentcached.NewCommentModel(c.MongoConf.Source, c.MongoConf.DataBase, c.MongoConf.CollComment, c.RedisConf),
		HistoryModel: commenthistory.NewHistoryModel(c.MongoConf.Source, c.MongoConf.DataBase, c.MongoConf.CollHistory),
	}
}
