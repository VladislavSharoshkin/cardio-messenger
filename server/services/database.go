package services

import (
	"awesomeProject/database"
	"awesomeProject/internal"
	"awesomeProject/utils"
	. "github.com/go-jet/jet/v2/postgres"
)

func GetBranches() map[string]interface{} {
	stmt := SELECT(database.BranchSelect()).
		FROM(database.BranchFrom())

	var branches []database.Branch
	err := stmt.Query(database.JetDB, &branches)
	if err != nil {
		return internal.LogError(err.Error(), 147947)
	}

	res := utils.Message(true, "")
	res["Branches"] = branches
	return res
}

func GetStaffs() map[string]interface{} {
	stmt := SELECT(database.StaffSelect()).
		FROM(database.StaffFrom())

	var staffs []database.Staff
	err := stmt.Query(database.JetDB, &staffs)
	if err != nil {
		return internal.LogError(err.Error(), 158636)
	}

	res := utils.Message(true, "")
	res["Staffs"] = staffs
	return res
}
