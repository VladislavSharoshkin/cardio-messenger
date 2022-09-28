package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	"github.com/go-jet/jet/v2/qrm"
)

type ActivityType int64

const (
	ActivityTypeStop ActivityType = iota + 1
	ActivityTypeTyping
)

type Activity struct {
	model.Activitys
}

var activityInsertedColumns = Activitys.MutableColumns.Except(Activitys.CreatedAt, Activitys.DeletedAt, Activitys.UpdatedAt)

func InsertActivity(db qrm.DB, activity model.Activitys) (model.Activitys, error) {

	stmt := Activitys.INSERT(activityInsertedColumns).
		MODEL(activity).RETURNING(Activitys.AllColumns)

	if err := stmt.QueryContext(Ctx, db, &activity); err != nil {
		return activity, err
	}

	return activity, nil
}
