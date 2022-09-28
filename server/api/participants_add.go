package api

type ParticipantsAdd struct {
	SenderId int64
	ChatId   int64
	Ids      []int64
}