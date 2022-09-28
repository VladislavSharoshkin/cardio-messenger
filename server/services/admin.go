package services

import (
	"awesomeProject/api"
	"awesomeProject/crypto"
	"awesomeProject/database"
	"awesomeProject/internal"
	"awesomeProject/utils"
)

func UserAdd(reg api.Reg) map[string]interface{} {
	err := reg.Validate()
	if err != nil {
		return internal.LogError(err.Error(), 812368)
	}

	hashPass := crypto.PasswordHash(reg.Pass, crypto.RandomBytes256())
	user := database.UserInit(reg.Login, &hashPass, reg.FirstName, reg.LastName, nil, nil, nil)

	err = database.InsertUser(database.JetDB, &user)
	if err != nil {
		return internal.LogError(err.Error(), 296358)
	}

	res := utils.Message(true, "")
	res["User"] = user
	return res
}
