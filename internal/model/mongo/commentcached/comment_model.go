package commentcached

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

const CommentCollectionName = "comment"

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		FindByAuthorIdAndType(ctx context.Context, authorId string, _type string, skip int64, limit int64) ([]Comment, error)
		FindByParent(ctx context.Context, _type string, parentId string, skip int64, limit int64) ([]Comment, error)
		FindByReplyToAndType(ctx context.Context, _type string, replyTo string, skip int64, limit int64) ([]Comment, error)
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

func (c customCommentModel) FindByReplyToAndType(ctx context.Context, _type string, replyTo string, skip int64, limit int64) ([]Comment, error) {
	var data []Comment
	err := c.conn.Find(ctx, &data,
		bson.M{
			"replyTo": replyTo,
			"type":    _type,
		},
		&mopt.FindOptions{
			Skip:  &skip,
			Limit: &limit,
		})
	return data, err
}

func (c customCommentModel) FindByParent(ctx context.Context, _type string, parentId string, skip int64, limit int64) ([]Comment, error) {
	var data []Comment
	err := c.conn.Find(ctx, &data,
		bson.M{
			"parentId": parentId,
			"type":     _type,
		},
		&mopt.FindOptions{
			Skip:  &skip,
			Limit: &limit,
		})
	return data, err
}

func (c customCommentModel) FindByAuthorIdAndType(ctx context.Context, authorId string, _type string, skip int64, limit int64) ([]Comment, error) {
	var data []Comment
	err := c.conn.Find(ctx, &data,
		bson.M{
			"authorId": authorId,
			"type":     _type,
		},
		&mopt.FindOptions{
			Skip:  &skip,
			Limit: &limit,
		})
	return data, err
}

// NewCommentModel returns a model for the mongo.
func NewCommentModel(url, db string, c cache.CacheConf) CommentModel {
	conn := monc.MustNewModel(url, db, CommentCollectionName, c)
	return &customCommentModel{
		defaultCommentModel: newDefaultCommentModel(conn),
	}
}
