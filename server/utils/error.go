package utils

const (
	ErrorCodeBadToken int = iota + 1
	ErrorCodeDatabase
	ErrorCodeDecodeRequest
	ErrorCodeNotFound
	ErrorCodeNotValidation
	ErrorCodeBadWord
	ErrorCodeHZ
)

type Error struct {
	Code int
	Text string
	Key int64
}

func ErrorInit(code int, text string, key int64) Error {
	return Error{Code: code, Text: text, Key: key}
}

func (thisObject Error) ToResponse() map[string]interface{} {
	return map[string]interface{} {"Error" : thisObject}
}