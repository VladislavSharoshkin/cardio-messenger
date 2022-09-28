package api

import (
	"awesomeProject/internal"
	u "awesomeProject/utils"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DeleteMessage struct {
	SenderId int64
	ChatId   int64
	Ids      []int64
}

func (thisObject DeleteMessage) Validate() (bool, map[string]interface{}) {
	err := validation.ValidateStruct(&thisObject,
		validation.Field(&thisObject.ChatId, validation.Required),
		validation.Field(&thisObject.Ids, validation.Required),
	)
	if err != nil{
		return false, internal.LogError(err.Error(), 812425)
	}
	return true, u.Message(true, "Success")
}
