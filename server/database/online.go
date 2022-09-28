package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	"github.com/go-jet/jet/v2/qrm"
)

const (
	OnlineTypeOffline int64 = iota + 1
	OnlineTypeOnline
)

type Online struct {
	model.Onlines
}

var onlineInsertedColumns = Onlines.MutableColumns.Except(Onlines.CreatedAt, Onlines.DeletedAt, Onlines.UpdatedAt)

func OnlineInit(UserID int64, Type int64) model.Onlines {
	return model.Onlines{
		Type:   Type,
		UserID: UserID,
	}
}

func InsertOnline(db qrm.DB, online *model.Onlines) error {

	stmt := Onlines.INSERT(onlineInsertedColumns).
		MODEL(online).RETURNING(Onlines.AllColumns)

	if err := stmt.QueryContext(Ctx, db, online); err != nil {
		return err
	}

	return nil
}
