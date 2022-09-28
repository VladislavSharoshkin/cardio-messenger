package api

import (
	"awesomeProject/internal"
	u "awesomeProject/utils"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type SendMessage struct {
	SenderId int64
	Text     *string
	ChatId      int64
	Attachments []int64
	ForwardMessages []int64
}

func (thisObject SendMessage) Validate() (bool, map[string]interface{}) {
	err := validation.ValidateStruct(&thisObject,
		validation.Field(&thisObject.ChatId, validation.Required),
		//validation.Field(&thisObject.PushToken, validation.Required),
	)
	if err != nil{
		return false, internal.LogError(err.Error(), 249218)
	}
	return true, u.Message(true, "Success")
}
