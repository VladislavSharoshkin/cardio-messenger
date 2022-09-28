package api

import (
	"awesomeProject/internal"
	u "awesomeProject/utils"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ChangePassword struct {
	SenderId int64
	OldPassword string
	NewPassword string
}

func (thisObject ChangePassword) Validate() (bool, map[string]interface{}) {
	err := validation.ValidateStruct(&thisObject,
		validation.Field(&thisObject.OldPassword, validation.Required),
		validation.Field(&thisObject.NewPassword, validation.Required, validation.Length(5, 100)),
	)
	if err != nil{
		return false, internal.LogError(err.Error(), 842184)
	}
	return true, u.Message(true, "Success")
}
