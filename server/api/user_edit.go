package api

type UserEdit struct {
	SenderId int64
	FirstName    *string
	MiddleName   *string
	LastName     *string
	AvatarID *int64
}
