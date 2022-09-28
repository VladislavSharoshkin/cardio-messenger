package api

type GetMessages struct {
	ChatId int64
	SenderId int64
	StartMessageId *int64
	Count *int64
}