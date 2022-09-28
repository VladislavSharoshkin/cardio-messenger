package services

import (
	"awesomeProject/database"
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	"awesomeProject/internal"
	u "awesomeProject/utils"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/openlyinc/pointy"
	"strconv"
	"time"
)

func branchCompare(remoteList []model.Branches, localList []model.Branches) ([]model.Branches, []model.Branches, []model.Branches) {
	var forAdd []model.Branches
	var forUpdate []model.Branches
	for _, remote := range remoteList {
		if ok, _ := database.BranchContainsRId(remote.SyncRemoteID, localList); ok {
			//if !local.Compare(remote) {
			//	local.Edit(remote)
			//	forUpdate = append(forUpdate, local)
			//}
		} else {
			forAdd = append(forAdd, remote)
		}
	}

	var forDelete []model.Branches
	//for _, local := range localList {
	//	if ok, _ := BranchContainsRId(local.SyncRemoteId, remoteList); !ok {
	//		forDelete = append(forDelete, local)
	//	}
	//}
	return forAdd, forUpdate, forDelete
}

func userCompare(remoteList []model.Users, localList []model.Users) ([]model.Users, []model.Users, []model.Users) {
	var forAdd []model.Users
	var forUpdate []model.Users
	for _, remote := range remoteList {
		if ok, _ := database.UserContainsRId(remote.SyncRemoteID, localList); ok {
			//if !local.Compare(remote) {
			//	local.Edit(remote)
			//	forUpdate = append(forUpdate, local)
			//}
		} else {
			forAdd = append(forAdd, remote)
		}
	}

	var forDelete []model.Users
	for _, local := range localList {
		if ok, _ := database.UserContainsRId(local.SyncRemoteID, remoteList); !ok {
			forDelete = append(forDelete, local)
		}
	}
	return forAdd, forUpdate, forDelete
}
func staffCompare(remoteList []model.Staffs, localList []model.Staffs) ([]model.Staffs, []model.Staffs, []model.Staffs) {
	var forAdd []model.Staffs
	var forUpdate []model.Staffs
	for _, remote := range remoteList {
		if ok, _ := database.StaffContainsRId(remote.SyncRemoteID, localList); ok {
			//if !local.Compare(remote) {
			//	local.Edit(remote)
			//	forUpdate = append(forUpdate, local)
			//}
		} else {
			forAdd = append(forAdd, remote)
		}
	}

	var forDelete []model.Staffs
	//for _, local := range localList {
	//	if ok, _ := BranchContainsRId(local.SyncRemoteId, remoteList); !ok {
	//		forDelete = append(forDelete, local)
	//	}
	//}
	return forAdd, forUpdate, forDelete
}

func userStaffCompare(remoteList []model.UserStaffs, localList []model.UserStaffs) ([]model.UserStaffs, []model.UserStaffs, []model.UserStaffs) {
	var forAdd []model.UserStaffs
	var forUpdate []model.UserStaffs
	for _, remote := range remoteList {
		if ok, _ := database.UserStaffContainsRId(remote.SyncRemoteID, localList); ok {
			//if !local.Compare(remote) {
			//	local.Edit(remote)
			//	forUpdate = append(forUpdate, local)
			//}
		} else {
			forAdd = append(forAdd, remote)
		}
	}

	var forDelete []model.UserStaffs
	//for _, local := range localList {
	//	if ok, _ := BranchContainsRId(local.SyncRemoteId, remoteList); !ok {
	//		forDelete = append(forDelete, local)
	//	}
	//}
	return forAdd, forUpdate, forDelete
}

//
//func userCompare (remoteList []User, localList []User) ([]User, []primitive.ObjectID, []database.ForUpdateMany) {
//	var forAdd []User
//	var forUpdate []database.ForUpdateMany
//	for _, remote := range remoteList {
//		if local := remote.ContainsRId(localList); local != nil {
//			if !local.Compare(remote) {
//				local.SyncEdit(remote)
//				forUpdate = append(forUpdate, database.ForUpdateManyInit(local.Id, local))
//			}
//		} else {
//			forAdd = append(forAdd, remote)
//		}
//	}
//
//	var forDelete []primitive.ObjectID
//	for _, local := range localList {
//		if local.ContainsRId(remoteList) == nil {
//			forDelete = append(forDelete, local.Id)
//		}
//	}
//	return forAdd, forDelete, forUpdate
//}
//
//func userStaffCompare (remoteList []UserStaff, localList []UserStaff) ([]UserStaff, []primitive.ObjectID, []database.ForUpdateMany) {
//	var forAdd []UserStaff
//	var forUpdate []database.ForUpdateMany
//	for _, remote := range remoteList {
//		if local := remote.ContainsRId(localList); local != nil {
//			if !local.Compare(remote) {
//				local.Edit(remote)
//				forUpdate = append(forUpdate, database.ForUpdateManyInit(local.Id, local))
//			}
//		} else {
//			forAdd = append(forAdd, remote)
//		}
//	}
//
//	var forDelete []primitive.ObjectID
//	for _, local := range localList {
//		if local.ContainsRId(remoteList) == nil {
//			forDelete = append(forDelete, local.Id)
//		}
//	}
//	return forAdd, forDelete, forUpdate
//}

func syncBranch(branchListRemote []model.Branches, sync database.Sync) bool {
	stmt := SELECT(database.BranchSelect()).
		FROM(database.BranchFrom()).
		WHERE(Branches.DeletedAt.GT_EQ(TimestampzT(time.Now())).
			AND(Branches.SyncID.EQ(Int(sync.ID))))

	var branchListLocal []model.Branches
	if err := stmt.Query(database.JetDB, &branchListLocal); err != nil {
		return false
	}

	forAdd, _, _ := branchCompare(branchListRemote, branchListLocal)
	database.InsertBranchs(database.JetDB, &forAdd)
	//DbDeleteMany(GetDbBranch(), forDel)
	//DbUpdateMany(GetDbBranch(), forUpd)

	return true
}

func syncStaffs(branchListRemote []model.Staffs, sync database.Sync) bool {
	stmt := SELECT(database.StaffSelect()).
		FROM(database.StaffFrom()).
		WHERE(Staffs.DeletedAt.GT_EQ(TimestampzT(time.Now())).
			AND(Staffs.SyncID.EQ(Int(sync.ID))))

	var branchListLocal []model.Staffs
	if err := stmt.Query(database.JetDB, &branchListLocal); err != nil {
		return false
	}

	forAdd, _, _ := staffCompare(branchListRemote, branchListLocal)
	database.InsertStaffs(database.JetDB, &forAdd)
	//DbDeleteMany(GetDbBranch(), forDel)
	//DbUpdateMany(GetDbBranch(), forUpd)

	return true
}

func syncUser(userListRemote []model.Users, sync database.Sync) bool {
	stmt := SELECT(database.UserSelect()).
		FROM(database.UserFrom()).
		WHERE(Users.DeletedAt.GT_EQ(TimestampzT(time.Now())).
			AND(Users.SyncID.EQ(Int(sync.ID))))

	var branchListLocal []model.Users
	if err := stmt.Query(database.JetDB, &branchListLocal); err != nil {
		return false
	}

	forAdd, _, forDel := userCompare(userListRemote, branchListLocal)
	database.InsertUsers(database.JetDB, &forAdd)
	database.DeleteUsers(database.JetDB, database.UserGetIDs(forDel))

	return true
}

//
//func syncStaff(staffListRemote []Staff, sync Sync) bool {
//	ok, staffListLocal, _ := DbStaffListGet(database.StaffGetQuery(database.MatchNotDeleted(database.And([]bson.M{
//		database.Eq("SyncId", sync.Id),
//	}))), false)
//	if !ok {
//		return false
//	}
//
//	forAdd, forDel, forUpd := staffCompare(staffListRemote, staffListLocal)
//	AddStaffList(forAdd)
//	DbDeleteMany(GetDbStaff(), forDel)
//	DbUpdateMany(GetDbStaff(), forUpd)
//
//	return true
//}
//
//func syncUserStaff(userStaffListRemote []UserStaff, sync Sync) bool {
//	ok, userStaffListLocal, _ := DbUserStaffListGet(database.UserStaffGetQuery(database.MatchNotDeleted(database.And([]bson.M{
//		database.Eq("SyncId", sync.Id),
//	}))), false)
//	if !ok {
//		return false
//	}
//
//	forAdd, forDel, forUpd := userStaffCompare(userStaffListRemote, userStaffListLocal)
//	AddUserStaffList(forAdd)
//	DbDeleteMany(GetDbUserStaff(), forDel)
//	DbUpdateMany(GetDbUserStaff(), forUpd)
//
//	return true
//}

func syncUserStaff(branchListRemote []model.UserStaffs, sync database.Sync) bool {
	stmt := SELECT(database.UserStaffSelect()).
		FROM(database.UserStaffFrom()).
		WHERE(UserStaffs.DeletedAt.GT_EQ(TimestampzT(time.Now())).
			AND(UserStaffs.SyncID.EQ(Int(sync.ID))))

	var branchListLocal []model.UserStaffs
	if err := stmt.Query(database.JetDB, &branchListLocal); err != nil {
		return false
	}

	forAdd, _, _ := userStaffCompare(branchListRemote, branchListLocal)
	database.InsertUserStaffs(database.JetDB, &forAdd)
	//DbDeleteMany(GetDbBranch(), forDel)
	//DbUpdateMany(GetDbBranch(), forUpd)

	return true
}

func getBranch(rows *sql.Rows, sync database.Sync) []model.Branches {
	var branchList []model.Branches
	for rows.Next() {
		var branch = model.Branches{}
		err := rows.Scan(&branch.SyncRemoteID, &branch.Name)
		if err != nil {
			return nil
		}
		branchList = append(branchList, database.BranchInit(branch.Name, &sync.ID, branch.SyncRemoteID))
	}
	return branchList
}

func getUser(rows *sql.Rows, sync database.Sync) []model.Users {
	var userList []model.Users
	for rows.Next() {
		var user = model.Users{}
		err := rows.Scan(&user.SyncRemoteID, &user.Login, &user.FirstName, &user.MiddleName, &user.LastName)
		if err != nil {
			return nil
		}
		userList = append(userList, database.UserInit(user.Login, nil, user.FirstName, user.LastName, user.MiddleName, &sync.ID, user.SyncRemoteID))
	}
	return userList
}

func getStaff(rows *sql.Rows, sync database.Sync) []model.Staffs {
	var staffList []model.Staffs
	for rows.Next() {
		var staff = model.Staffs{}
		err := rows.Scan(&staff.SyncRemoteID, &staff.Name)
		if err != nil {
			return nil
		}
		staffList = append(staffList, database.StaffInit(staff.Name, &sync.ID, staff.SyncRemoteID))
	}
	return staffList
}

func getUserStaff(rows *sql.Rows, sync database.Sync) []model.UserStaffs {

	stmt := SELECT(database.BranchSelect()).
		FROM(database.BranchFrom()).
		WHERE(Branches.DeletedAt.GT_EQ(TimestampzT(time.Now())))

	var branchListLocal []model.Branches
	if err := stmt.Query(database.JetDB, &branchListLocal); err != nil {
		return nil
	}

	stmt = SELECT(database.StaffSelect()).
		FROM(database.StaffFrom()).
		WHERE(Staffs.DeletedAt.GT_EQ(TimestampzT(time.Now())))

	var staffListLocal []model.Staffs
	if err := stmt.Query(database.JetDB, &staffListLocal); err != nil {
		return nil
	}

	stmt = SELECT(database.UserSelect()).
		FROM(database.UserFrom()).
		WHERE(Users.DeletedAt.GT_EQ(TimestampzT(time.Now())))

	var userListLocal []model.Users
	if err := stmt.Query(database.JetDB, &userListLocal); err != nil {
		return nil
	}

	var userStaffList []model.UserStaffs
	for rows.Next() {
		var userStaff = model.UserStaffs{}
		err := rows.Scan(&userStaff.SyncRemoteID, &userStaff.SyncRemoteUserID, &userStaff.SyncRemoteStaffID, &userStaff.SyncRemoteBranchID)
		if err != nil {
			return nil
		}
		_, user := database.UserContainsRId(userStaff.SyncRemoteUserID, userListLocal)
		_, staff := database.StaffContainsRId(userStaff.SyncRemoteStaffID, staffListLocal)
		_, branch := database.BranchContainsRId(userStaff.SyncRemoteBranchID, branchListLocal)
		if user.ID != 0 && staff.ID != 0 && branch.ID != 0 {
			userStaffList = append(userStaffList, database.UserStaffInit(
				user.ID, branch.ID, staff.ID, &sync.ID, userStaff.SyncRemoteID, userStaff.SyncRemoteUserID,
				userStaff.SyncRemoteBranchID, userStaff.SyncRemoteStaffID))
		}
	}
	return userStaffList
}

func syncGeneralChat(creatorId int64) error {
	stmt := SELECT(database.ChatSelect(creatorId)).
		FROM(database.ChatFrom(creatorId)).
		WHERE(Chats.NameUnique.EQ(String(database.ChatNameUniqueGeneral)))

	var chat database.Chat
	err := stmt.Query(database.JetDB, &chat)
	if err != nil && err != qrm.ErrNoRows {
		return err
	}

	stmt = SELECT(Users.ID).
		FROM(Users).
		WHERE(database.NotDeleted(Users.DeletedAt))
	var userIDs []int64
	err = stmt.Query(database.JetDB, &userIDs)
	if err != nil && err != qrm.ErrNoRows {
		return err
	}

	forAdd, _ := u.CompareSlice(database.ParticipantsGetUserIDs(chat.Participants), userIDs)

	if chat.ID == 0 {
		database.InsertChat(database.JetDB, database.ChatInit(database.ChatTypeGroup, creatorId, pointy.String("Общая группа")), forAdd)
	} else {
		database.InsertParticipants(database.JetDB, chat.ID, forAdd)
	}

	return nil
}

func syncBranchChat(creatorId int64, branchID int64) error {
	stmt := SELECT(database.ChatSelect(creatorId)).
		FROM(database.ChatFrom(creatorId)).
		WHERE(Chats.NameUnique.EQ(String(database.ChatNameUniqueBranch + "_" + strconv.FormatInt(branchID, 10))))

	var chat database.Chat
	err := stmt.Query(database.JetDB, &chat)
	if err != nil && err != qrm.ErrNoRows {
		return err
	}

	stmt = SELECT(Users.ID).
		FROM(Users).
		WHERE(database.NotDeleted(Users.DeletedAt))
	var userIDs []int64
	err = stmt.Query(database.JetDB, &userIDs)
	if err != nil && err != qrm.ErrNoRows {
		return err
	}

	forAdd, _ := u.CompareSlice(database.ParticipantsGetUserIDs(chat.Participants), userIDs)

	if chat.ID == 0 {
		database.InsertChat(database.JetDB, database.ChatInit(database.ChatTypeGroup, creatorId, pointy.String("Общая группа")), forAdd)
	} else {
		database.InsertParticipants(database.JetDB, chat.ID, forAdd)
	}

	return nil
}

func start(db *sql.DB, sync database.Sync) bool {
	rows, err := db.Query(sync.Query)
	if err != nil {
		return false
	}

	switch sync.Type {
	case database.TypeBranch:
		syncBranch(getBranch(rows, sync), sync)
	case database.TypeUser:
		syncUser(getUser(rows, sync), sync)
	case database.TypeStaff:
		syncStaffs(getStaff(rows, sync), sync)
	case database.TypeUserStaff:
		syncUserStaff(getUserStaff(rows, sync), sync)
	default:
		return false
	}

	return true
}

func StartAll() map[string]interface{} {
	stmt := SELECT(database.SyncSelect()).
		FROM(database.SyncFrom()).
		WHERE(Syncs.DeletedAt.GT_EQ(TimestampzT(time.Now())))

	var syncs []database.Sync
	if err := stmt.Query(database.JetDB, &syncs); err != nil {
		return internal.LogError(err.Error(), 293195)
	}

	var connString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", "10.30.10.33", "LPU\\SharoshkinVM", "Shvm157", 1433, "MedIS")
	db, err := sql.Open("mssql", connString)
	if err != nil {
		return internal.LogError(err.Error(), 293593)
	}
	defer db.Close()

	for _, sync := range syncs {
		start(db, sync)
	}

	syncGeneralChat(3326)

	return u.Message(true, "StartAll")
}
