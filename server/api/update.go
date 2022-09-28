package api

const (
	UpdateTypeUnknown UpdateType = iota // 0
	UpdateTypeAdd
	UpdateTypeDelete
	UpdateTypeEdit
)

type UpdateType int64

const (
	UpdateDataTypeUnknown UpdateDataType = iota // 0
	UpdateDataTypeChat
	UpdateDataTypeMessage
)

type UpdateDataType int64

type Update struct {
	UpdateType UpdateType
	DataType   UpdateDataType
	Data       interface{}
}

func UpdateInit(UpdateType UpdateType, DataType UpdateDataType, Data interface{}) Update {
	return Update{UpdateType: UpdateType, Data: Data, DataType: DataType}
}