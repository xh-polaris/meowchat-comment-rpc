package svc

import (
	"github.com/xh-polaris/meowchat-comment-rpc/internal/config"
	"github.com/xh-polaris/meowchat-comment-rpc/internal/model/mongo/commentcached"
	"github.com/xh-polaris/meowchat-comment-rpc/internal/model/mongo/commenthistory"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel commentcached.CommentModel
	HistoryModel commenthistory.HistoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		CommentModel: commentcached.NewCommentModel(c.Mongo.URL, c.Mongo.DB, c.Cache),
		HistoryModel: commenthistory.NewHistoryModel(c.Mongo.URL, c.Mongo.DB),
	}
}
