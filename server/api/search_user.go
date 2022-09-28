package api

import (
	"awesomeProject/internal"
	u "awesomeProject/utils"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type SearchUsers struct {
	SearchString string
	Branches []int64
	Staffs []int64
}

func (thisObject SearchUsers) Validate() (bool, map[string]interface{}) {
	err := validation.ValidateStruct(&thisObject,
		validation.Field(&thisObject.SearchString, validation.Required, validation.Length(3,1000)),
	)
	if err != nil{
		return false, internal.LogError(err.Error(), 559493)
	}
	return true, u.Message(true, "Success")
}
