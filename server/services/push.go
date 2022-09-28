package services

import (
	"awesomeProject/api"
	"awesomeProject/crypto"
	"awesomeProject/database"
	"awesomeProject/internal"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	. "awesomeProject/gen/postgres/public/table"

	"awesomeProject/utils"
	. "github.com/go-jet/jet/v2/postgres"
)

func PushAdd(pushAdd api.PushAdd) map[string]interface{} {
	serverPush := database.ServerPushInit(pushAdd.Push, crypto.NewToken())
	err := database.InsertServerPush(database.JetDB, &serverPush)
	if err != nil {
		return internal.LogError(err.Error(), 104651)
	}

	response := utils.Message(true, "")
	response["ServerPush"] = serverPush
	return response
}

func PushSend(pushSend api.PushSend) map[string]interface{} {

	stmt := SELECT(
		ServerPushs.Push,
	).FROM(
		ServerPushs,
	).WHERE(
		database.NotDeleted(ServerPushs.DeletedAt).
			AND(ServerPushs.Token.IN(database.InString(pushSend.Tokens)...)),
	).DISTINCT(ServerPushs.Push)
	var pushs []string
	err := stmt.Query(database.JetDB, &pushs)
	if err != nil {
		return internal.LogError(err.Error(), 893274)
	}

	err = internal.SendPushNotifications(pushs, pushSend.Title, pushSend.Body)
	if err != nil {
		return internal.LogError(err.Error(), 248347)
	}

	response := utils.Message(true, "")
	return response
}

func PushSend2(userIDs []int64, title string, body string) error {
	stmt := SELECT(
		Tokens.Push,
	).FROM(
		Tokens,
	).WHERE(
		database.NotDeleted(Tokens.DeletedAt).
			AND(Tokens.UserID.IN(database.InInt(userIDs)...)),
	).DISTINCT(Tokens.Push)
	var tokens []string
	err := stmt.Query(database.JetDB, &tokens)
	if err != nil {
		return err
	}

	data := api.PushSendInit(tokens, title, body)
	json_data, err := json.Marshal(data)

	if err != nil {
		return err
	}

	skipTlsClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		//Timeout: 10,
	}

	resp, err := skipTlsClient.Post("https://"+utils.Settings.PushServer+"/push/send", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	return nil
}
