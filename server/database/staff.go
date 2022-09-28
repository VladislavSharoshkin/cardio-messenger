package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/openlyinc/pointy"
)

type Staff struct {
	model.Staffs
}

var staffInsertedColumns = Staffs.MutableColumns.Except(Staffs.CreatedAt, Staffs.DeletedAt, Staffs.UpdatedAt)

func StaffInit(Name string, SyncID *int64, SyncRemoteID *string) model.Staffs {
	return model.Staffs{
		Name:         Name,
		SyncID:       SyncID,
		SyncRemoteID: SyncRemoteID,
	}
}

func StaffContainsRId(SyncRemoteId *string, roleList []model.Staffs) (bool, model.Staffs) {
	for _, element := range roleList {
		if pointy.StringValue(element.SyncRemoteID, "") == pointy.StringValue(SyncRemoteId, "") {
			return true, element
		}
	}
	return false, model.Staffs{}
}

func StaffSelect() ProjectionList {
	return ProjectionList{
		Staffs.AllColumns,
	}
}

func StaffFrom() ReadableTable {
	return Staffs
}

func InsertStaffs(db qrm.DB, staffs *[]model.Staffs) error {
	stmt := Staffs.INSERT(staffInsertedColumns).
		MODELS(staffs).RETURNING(Staffs.AllColumns)

	if err := stmt.QueryContext(Ctx, db, staffs); err != nil {
		return err
	}
	return nil
}
