//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Chats struct {
	ID         int64 `sql:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	CreatorID  int64
	Type       int64
	Name       *string
	TypeUnique *int64
	AvatarID   *int64
	About      *string
	NameUnique *string
}