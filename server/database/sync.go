package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	. "github.com/go-jet/jet/v2/postgres"
)

const (
	TypeBranch int64 = iota + 1
	TypeUser
	TypeStaff
	TypeUserStaff
)

type Sync struct {
	model.Syncs
}

var syncInsertedColumns = ForwardMessages.MutableColumns.Except(Syncs.CreatedAt, Syncs.DeletedAt, Syncs.UpdatedAt)

func syncInit(MessageID int64, ForwardMessageID int64) model.ForwardMessages {
	return model.ForwardMessages{
		MessageID:        MessageID,
		ForwardMessageID: ForwardMessageID,
	}
}

func SyncSelect() ProjectionList {
	return ProjectionList{
		Syncs.AllColumns,
	}
}

func SyncFrom() ReadableTable {
	return Syncs
}
