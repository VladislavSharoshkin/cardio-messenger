package services

import (
	"awesomeProject/database"
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	"awesomeProject/internal"
	"awesomeProject/utils"
	. "github.com/go-jet/jet/v2/postgres"
	"io/ioutil"
	"strconv"
)

func FileUpload(file model.Files) map[string]interface{} {

	err := database.InsertFile(database.JetDB, &file)
	if err != nil {
		return internal.LogError(err.Error(), 167427)
	}

	resp := utils.Message(true, "")
	resp["File"] = file
	return resp
}

func FileDownload(token string) (map[string]interface{}, []byte) {

	fileId, _ := strconv.ParseInt(token, 10, 64)

	stmt := SELECT(
		database.FileSelect(),
	).FROM(database.FileFrom()).WHERE(
		Files.Token.EQ(String(token)).
			OR(Files.ID.EQ(Int(fileId))),
	)

	var file database.File
	err := stmt.Query(database.JetDB, &file)
	if err != nil {
		return internal.LogError(err.Error(), 622993), nil
	}

	fileBytes, err := ioutil.ReadFile("files/" + file.Hash)
	if err != nil {
		return internal.LogError(err.Error(), 135935), nil
	}

	resp := utils.Message(true, "")
	return resp, fileBytes
}
