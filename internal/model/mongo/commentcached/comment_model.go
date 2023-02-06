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
		FindByAuthorIdAndType(ctx context.Context, authorId string, _type string, skip int64, limit int64) ([]Comment, int64, error)
		FindByParent(ctx context.Context, _type string, parentId string, skip int64, limit int64) ([]Comment, int64, error)
		FindByReplyToAndType(ctx context.Context, _type string, replyTo string, skip int64, limit int64) ([]Comment, int64, error)
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

func (c customCommentModel) FindByReplyToAndType(ctx context.Context, _type string, replyTo string, skip int64, limit int64) ([]Comment, int64, error) {
	var data []Comment
	if err := c.conn.Find(ctx, &data,
		bson.M{
			"replyTo": replyTo,
			"type":    _type,
		},
		&mopt.FindOptions{
			Skip:  &skip,
			Limit: &limit,
			Sort:  bson.M{"createAt": -1},
		}); err != nil {
		return nil, 0, err
	}
	count, err := c.conn.CountDocuments(ctx, bson.M{"replyTo": replyTo,
		"type": _type})
	if err != nil {
		return nil, 0, err
	}
	return data, count, err
}

func (c customCommentModel) FindByParent(ctx context.Context, _type string, parentId string, skip int64, limit int64) ([]Comment, int64, error) {
	var data []Comment
	if err := c.conn.Find(ctx, &data,
		bson.M{
			"parentId": parentId,
			"type":     _type,
		},
		&mopt.FindOptions{
			Skip:  &skip,
			Limit: &limit,
			Sort:  bson.M{"createAt": -1},
		}); err != nil {
		return nil, 0, err
	}
	count, err := c.conn.CountDocuments(ctx, bson.M{
		"parentId": parentId,
		"type":     _type,
	})
	if err != nil {
		return nil, 0, err
	}
	return data, count, err
}

func (c customCommentModel) FindByAuthorIdAndType(ctx context.Context, authorId string, _type string, skip int64, limit int64) ([]Comment, int64, error) {
	var data []Comment
	if err := c.conn.Find(ctx, &data,
		bson.M{
			"authorId": authorId,
			"type":     _type,
		},
		&mopt.FindOptions{
			Skip:  &skip,
			Limit: &limit,
			Sort:  bson.M{"createAt": -1},
		}); err != nil {
		return nil, 0, err
	}
	count, err := c.conn.CountDocuments(ctx, bson.M{
		"authorId": authorId,
		"type":     _type,
	})
	if err != nil {
		return nil, 0, err
	}
	return data, count, err
}

// NewCommentModel returns a model for the mongo.
func NewCommentModel(url, db string, c cache.CacheConf) CommentModel {
	conn := monc.MustNewModel(url, db, CommentCollectionName, c)
	return &customCommentModel{
		defaultCommentModel: newDefaultCommentModel(conn),
	}
}
