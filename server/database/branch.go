package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/openlyinc/pointy"
)

type Branch struct {
	model.Branches
}

var branchInsertedColumns = Branches.MutableColumns.Except(Branches.CreatedAt, Branches.DeletedAt, Branches.UpdatedAt)

func BranchInit(Name string, SyncID *int64, SyncRemoteID *string) model.Branches {
	return model.Branches{
		Name:         Name,
		SyncID:       SyncID,
		SyncRemoteID: SyncRemoteID,
	}
}

func BranchContainsRId(SyncRemoteId *string, roleList []model.Branches) (bool, model.Branches) {
	for _, element := range roleList {
		if pointy.StringValue(element.SyncRemoteID, "") == pointy.StringValue(SyncRemoteId, "") {
			return true, element
		}
	}
	return false, model.Branches{}
}

func BranchSelect() ProjectionList {
	return ProjectionList{
		Branches.AllColumns,
	}
}

func BranchFrom() ReadableTable {
	return Branches
}

func InsertBranchs(db qrm.DB, branchs *[]model.Branches) error {
	stmt := Branches.INSERT(branchInsertedColumns).
		MODELS(branchs).RETURNING(Branches.AllColumns)

	if err := stmt.QueryContext(Ctx, db, branchs); err != nil {
		return err
	}
	return nil
}
