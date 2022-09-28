package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type Read struct {
	model.Reads
}

var readInsertedColumns = Reads.MutableColumns.Except(Reads.CreatedAt, Reads.DeletedAt, Reads.UpdatedAt)

func ReadInit(MessageID int64, UserID int64) model.Reads {
	return model.Reads{
		MessageID: MessageID,
		UserID:    UserID,
	}
}

func ReadSelect() ProjectionList {
	return ProjectionList{
		Reads.AllColumns,
	}
}

func ReadFrom() ReadableTable {
	return Reads
}

func InsertRead(db qrm.DB, read *model.Reads) error {
	stmt := Reads.INSERT(readInsertedColumns).
		MODEL(read).RETURNING(Reads.AllColumns)

	if err := stmt.QueryContext(Ctx, db, read); err != nil {
		return err
	}
	return nil
}
