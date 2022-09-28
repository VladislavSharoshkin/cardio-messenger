package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type ServerPush struct {
	model.ServerPushs
}

var serverPushInsertedColumns = ServerPushs.MutableColumns.Except(ServerPushs.CreatedAt, ServerPushs.DeletedAt, ServerPushs.UpdatedAt)

func ServerPushInit(Push string, Token string) model.ServerPushs {
	return model.ServerPushs{
		Push:  Push,
		Token: Token,
	}
}

func ServerPushSelect() ProjectionList {
	return ProjectionList{
		ServerPushs.AllColumns,
	}
}

func ServerPushFrom() ReadableTable {
	return ServerPushs
}

func InsertServerPush(db qrm.DB, serverPush *model.ServerPushs) error {
	stmt := ServerPushs.INSERT(serverPushInsertedColumns).
		MODEL(serverPush).RETURNING(ServerPushs.AllColumns)

	if err := stmt.QueryContext(Ctx, db, serverPush); err != nil {
		return err
	}
	return nil
}
