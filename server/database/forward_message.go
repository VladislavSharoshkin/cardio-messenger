package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type ForwardMessage struct {
	model.ForwardMessages
}

var forwardMessageInsertedColumns = ForwardMessages.MutableColumns.Except(ForwardMessages.CreatedAt, ForwardMessages.DeletedAt, ForwardMessages.UpdatedAt)

func forwardMessageInit(MessageID int64, ForwardMessageID int64) model.ForwardMessages {
	return model.ForwardMessages{
		MessageID:        MessageID,
		ForwardMessageID: ForwardMessageID,
	}
}

func ForwardMessageSelect() ProjectionList {
	return ProjectionList{
		ForwardMessages.AllColumns,
	}
}

func ForwardMessageFrom() ReadableTable {
	return ForwardMessages
}

func InsertForwardMessages(db qrm.DB, messageID int64, forwardIDs []int64) ([]model.ForwardMessages, error) {
	var forwardMessages []model.ForwardMessages
	for _, forwardID := range forwardIDs {
		forwardMessages = append(forwardMessages, forwardMessageInit(messageID, forwardID))
	}

	stmt := ForwardMessages.INSERT(forwardMessageInsertedColumns).
		MODELS(forwardMessages).RETURNING()

	if err := stmt.QueryContext(Ctx, db, &forwardMessages); err != nil {
		return nil, err
	}
	return forwardMessages, nil
}
