//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type UserStaffs struct {
	ID                 int64 `sql:"primary_key"`
	UserID             int64
	BranchID           int64
	StaffID            int64
	SyncID             *int64
	SyncRemoteID       *string
	SyncRemoteUserID   *string
	SyncRemoteBranchID *string
	SyncRemoteStaffID  *string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          time.Time
}