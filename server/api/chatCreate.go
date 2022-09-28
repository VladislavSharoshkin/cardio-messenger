package api

import (
	"awesomeProject/internal"
	u "awesomeProject/utils"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateChat struct {
	SenderId int64
	Type int64
	Name *string
	Participants []int64
}

func (thisObject CreateChat) Validate() (bool, map[string]interface{}) {
	err := validation.ValidateStruct(&thisObject,
		validation.Field(&thisObject.Type, validation.Required),
	)
	if err != nil{
		return false, internal.LogError(err.Error(), 383434)
	}
	return true, u.Message(true, "Success")
}
