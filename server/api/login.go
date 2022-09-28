package api

import (
	"awesomeProject/internal"
	u "awesomeProject/utils"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Login struct {
	Login     string
	Pass      string
	Push *string
}

func (thisObject Login) Validate() (bool, map[string]interface{}) {
	err := validation.ValidateStruct(&thisObject,
		validation.Field(&thisObject.Login, validation.Required),
		validation.Field(&thisObject.Pass, validation.Required),
		//validation.Field(&thisObject.PushToken, validation.Required),
	)
	if err != nil{
		return false, internal.LogError(err.Error(), 915603)
	}
	return true, u.Message(true, "Success")
}
