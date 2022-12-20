package commenthistory

import "github.com/zeromicro/go-zero/core/stores/mon"

const HistoryCollectionName = "history"

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
func NewHistoryModel(url, db string) HistoryModel {
	conn := mon.MustNewModel(url, db, HistoryCollectionName)
	return &customHistoryModel{
		defaultHistoryModel: newDefaultHistoryModel(conn),
	}
}
