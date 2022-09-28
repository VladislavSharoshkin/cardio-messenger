package api

type ChatEdit struct {
	SenderId int64
	ChatID int64
	Name    *string
	About   *string
	AvatarID *int64
}
