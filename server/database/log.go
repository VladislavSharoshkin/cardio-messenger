package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

const (
	LogTypeInfo int64 = iota + 1
	LogTypeWarning
	LogTypeError
	LogTypeLoginFail
)

type Log struct {
	model.Logs
}

var logInsertedColumns = Logs.MutableColumns.Except(Logs.CreatedAt, Logs.DeletedAt, Logs.UpdatedAt)

func LogInit(Type int64, Text string, RemoteAddr *string, RequestURI *string, ErrorKey *int64) model.Logs {
	return model.Logs{
		Type:       Type,
		Text:       Text,
		RemoteAddr: RemoteAddr,
		RequestURI: RequestURI,
		ErrorKey:   ErrorKey,
	}
}

func LogSelect() ProjectionList {
	return ProjectionList{
		Logs.AllColumns,
	}
}

func LogFrom() ReadableTable {
	return Logs
}

func LogInsert(db qrm.DB, log *model.Logs) error {
	stmt := Logs.INSERT(logInsertedColumns).
		MODEL(log).
		RETURNING(Logs.AllColumns)
	if err := stmt.QueryContext(Ctx, db, log); err != nil {
		return err
	}

	return nil
}
