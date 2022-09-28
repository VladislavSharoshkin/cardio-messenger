package api

import (
	"awesomeProject/internal"
	u "awesomeProject/utils"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ReadMessage struct {
	SenderId int64
	ChatId int64
	MessageId int64
}

func (thisObject ReadMessage) Validate() (bool, map[string]interface{}) {
	err := validation.ValidateStruct(&thisObject,
		validation.Field(&thisObject.ChatId, validation.Required),
		validation.Field(&thisObject.MessageId, validation.Required),
	)
	if err != nil{
		return false, internal.LogError(err.Error(), 522903)
	}
	return true, u.Message(true, "Success")
}
