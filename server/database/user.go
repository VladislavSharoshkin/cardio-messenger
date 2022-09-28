package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	"github.com/ahmetb/go-linq/v3"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/openlyinc/pointy"
	"strings"
)

type User struct {
	model.Users
	Avatar      string
	Subscribers []struct {
		model.Subscribers `alias:"UserSubscribers.*"`
		Subscriber        struct {
			model.Users `alias:"SubscribersSubscriber.*"`
		}
	}
	UserStaffs []struct {
		model.UserStaffs `alias:"UserUserStaffs.*"`
		Branch           struct {
			model.Branches `alias:"UserStaffsBranch.*"`
		}
		Staff struct {
			model.Staffs `alias:"UserStaffsStaff.*"`
		}
	}
	LastOnline *struct {
		model.Onlines `alias:"UserLastOnline.*"`
	}
}

var UserSubscribers = Subscribers.AS("UserSubscribers")
var SubscribersSubscriber = Users.AS("SubscribersSubscriber")
var UserUserStaffs = UserStaffs.AS("UserUserStaffs")
var UserStaffsBranch = Branches.AS("UserStaffsBranch")
var UserStaffsStaff = Staffs.AS("UserStaffsStaff")
var UserLastOnline = Onlines.AS("UserLastOnline")
var insertedColumns = Users.MutableColumns.Except(Users.CreatedAt, Users.DeletedAt, Users.UpdatedAt)

func UserInit(Login string, HashPass *string, FirstName string, LastName string, MiddleName *string, SyncID *int64,
	SyncRemoteID *string) model.Users {
	return model.Users{
		Login:        strings.ToLower(Login),
		HashPass:     HashPass,
		FirstName:    FirstName,
		LastName:     LastName,
		MiddleName:   MiddleName,
		SyncID:       SyncID,
		SyncRemoteID: SyncRemoteID,
	}
}

func UserGetIDs(users []model.Users) []int64 {
	var IDs []int64
	linq.From(users).Select(func(c interface{}) interface{} {
		return c.(model.Users).ID
	}).ToSlice(&IDs)
	return IDs
}

func LastOnlineLateral(user *UsersTable) SelectTable {
	return LATERAL(
		SELECT(
			UserLastOnline.AllColumns,
		).FROM(
			UserLastOnline,
		).WHERE(
			UserLastOnline.UserID.EQ(user.ID),
		).ORDER_BY(UserLastOnline.ID.DESC()).LIMIT(1),
	).AS("LastOnlineLateral")
}

func UserContainsRId(SyncRemoteId *string, roleList []model.Users) (bool, model.Users) {
	for _, element := range roleList {
		if pointy.StringValue(element.SyncRemoteID, "") == pointy.StringValue(SyncRemoteId, "") {
			return true, element
		}
	}
	return false, model.Users{}
}

func UserSelect() ProjectionList {
	return ProjectionList{
		Users.AllColumns, UserSubscribers.AllColumns, SubscribersSubscriber.AllColumns,
		String(GetFileUrl(1)).AS("User.Avatar"), UserUserStaffs.AllColumns, UserStaffsBranch.AllColumns,
		UserStaffsStaff.AllColumns, LastOnlineLateral(Users).AllColumns(),
		//CASE().WHEN(Chats.Type.EQ(Int(1))).THEN(companionAvatarId).
		//	ELSE(NULL).AS("Chat.AvatarID"),
	}
}

func UserFrom() ReadableTable {
	return Users.LEFT_JOIN(UserSubscribers, UserSubscribers.UserID.EQ(Users.ID)).
		LEFT_JOIN(SubscribersSubscriber, UserSubscribers.SubscriberID.EQ(SubscribersSubscriber.ID)).
		LEFT_JOIN(UserUserStaffs, UserUserStaffs.UserID.EQ(Users.ID)).
		LEFT_JOIN(UserStaffsBranch, UserStaffsBranch.ID.EQ(UserUserStaffs.BranchID)).
		LEFT_JOIN(UserStaffsStaff, UserStaffsStaff.ID.EQ(UserUserStaffs.StaffID)).
		LEFT_JOIN(LastOnlineLateral(Users), Bool(true))
}

func InsertUsers(db qrm.DB, users *[]model.Users) error {
	stmt := Users.INSERT(insertedColumns).
		MODELS(users).RETURNING(Users.AllColumns)

	if err := stmt.QueryContext(Ctx, db, users); err != nil {
		return err
	}
	return nil
}

func InsertUser(db qrm.DB, user *model.Users) error {
	stmt := Users.INSERT(insertedColumns).
		MODEL(user).RETURNING(Users.AllColumns)

	if err := stmt.QueryContext(Ctx, db, user); err != nil {
		return err
	}
	return nil
}

func UpdateUser(db qrm.DB, user *model.Users) error {
	stmt := Users.UPDATE(insertedColumns).
		MODEL(user).
		WHERE(Users.ID.EQ(Int(user.ID))).RETURNING(Users.AllColumns)

	err := stmt.QueryContext(Ctx, db, user)
	return err
}

func DeleteUsers(db qrm.DB, ids []int64) ([]model.Users, error) {
	updateStmt := Users.UPDATE(Users.DeletedAt).
		SET("now()").
		WHERE(Users.ID.IN(InInt(ids)...)).
		RETURNING(Users.AllColumns)

	var users []model.Users
	if err := updateStmt.QueryContext(Ctx, db, &users); err != nil {
		return nil, err
	}

	stmt := SELECT(Tokens.ID).
		FROM(Tokens).
		WHERE(Tokens.UserID.IN(InInt(UserGetIDs(users))...))

	var tokenIDs []int64
	err := stmt.QueryContext(Ctx, db, &tokenIDs)
	if err != nil {
		return nil, err
	}

	_, err = DeleteTokens(JetDB, tokenIDs)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func UserById(id int64) (User, error) {
	stmt := SELECT(
		UserSelect(),
	).FROM(
		UserFrom(),
	).WHERE(
		Users.ID.EQ(Int(id)),
	)
	var user User
	if err := stmt.Query(JetDB, &user); err != nil {
		return user, err
	}
	return user, nil
}
