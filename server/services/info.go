package services

import "awesomeProject/utils"

func Info() map[string]interface{} {
	res := utils.Message(true, "")
	res["Settings"] = utils.Settings.GetPublic()
	return res
}
