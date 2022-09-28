package api

type PushSend struct {
	Tokens []string
	Title string
	Body string
}

func PushSendInit(Tokens []string, Title string, Body string) PushSend {
	return PushSend{Tokens: Tokens, Title: Title, Body: Body}
}