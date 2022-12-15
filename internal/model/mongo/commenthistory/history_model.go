package commenthistory

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ HistoryModel = (*customHistoryModel)(nil)

type (
	// HistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHistoryModel.
	HistoryModel interface {
		historyModel
	}

	customHistoryModel struct {
		*defaultHistoryModel
	}
)

// NewHistoryModel returns a model for the mongo.
func NewHistoryModel(url, db, collection string) HistoryModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customHistoryModel{
		defaultHistoryModel: newDefaultHistoryModel(conn),
	}
}
