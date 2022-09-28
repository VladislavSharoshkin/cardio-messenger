package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/openlyinc/pointy"
)

type UserStaff struct {
	model.Staffs
}

var userStaffInsertedColumns = UserStaffs.MutableColumns.Except(UserStaffs.CreatedAt, UserStaffs.DeletedAt, UserStaffs.UpdatedAt)

func UserStaffInit(UserID int64, BranchID int64, StaffID int64, SyncID *int64, SyncRemoteID *string,
	SyncRemoteUserID *string, SyncRemoteBranchID *string, SyncRemoteStaffID *string) model.UserStaffs {
	return model.UserStaffs{
		UserID:             UserID,
		BranchID:           BranchID,
		StaffID:            StaffID,
		SyncID:             SyncID,
		SyncRemoteID:       SyncRemoteID,
		SyncRemoteUserID:   SyncRemoteUserID,
		SyncRemoteBranchID: SyncRemoteBranchID,
		SyncRemoteStaffID:  SyncRemoteStaffID,
	}
}

func UserStaffContainsRId(SyncRemoteId *string, roleList []model.UserStaffs) (bool, model.UserStaffs) {
	for _, element := range roleList {
		if pointy.StringValue(element.SyncRemoteID, "") == pointy.StringValue(SyncRemoteId, "") {
			return true, element
		}
	}
	return false, model.UserStaffs{}
}

func UserStaffSelect() ProjectionList {
	return ProjectionList{
		UserStaffs.AllColumns,
	}
}

func UserStaffFrom() ReadableTable {
	return UserStaffs
}

func InsertUserStaffs(db qrm.DB, userStaffs *[]model.UserStaffs) error {
	stmt := UserStaffs.INSERT(userStaffInsertedColumns).
		MODELS(userStaffs).RETURNING(UserStaffs.AllColumns)

	if err := stmt.QueryContext(Ctx, db, userStaffs); err != nil {
		return err
	}
	return nil
}
