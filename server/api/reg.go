package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Reg struct {
	Login string
	Pass string
	FirstName string
	LastName string
}

func (thisObject Reg) Validate() error {
	err := validation.ValidateStruct(&thisObject,
		validation.Field(&thisObject.Login, validation.Required, validation.Length(4, 30)),
		validation.Field(&thisObject.Pass, validation.Required, validation.Length(4, 30)),
	)
	return err
}
