package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type MessageFile struct {
	model.MessageFiles
}

var messageFileInsertedColumns = MessageFiles.MutableColumns.Except(MessageFiles.CreatedAt, MessageFiles.DeletedAt, MessageFiles.UpdatedAt)

func MessageFileInit(MessageID int64, FileID int64) model.MessageFiles {
	return model.MessageFiles{
		MessageID: MessageID,
		FileID:    FileID,
	}
}

func MessageFileSelect() ProjectionList {
	return ProjectionList{
		MessageFiles.AllColumns,
	}
}

func MessageFileFrom() ReadableTable {
	return MessageFiles
}

func InsertMessageFiles(db qrm.DB, messageID int64, fileIDs []int64) ([]model.MessageFiles, error) {
	var messageFiles []model.MessageFiles
	for _, fileID := range fileIDs {
		messageFiles = append(messageFiles, MessageFileInit(messageID, fileID))
	}

	stmt := MessageFiles.INSERT(messageFileInsertedColumns).
		MODELS(messageFiles).RETURNING()

	if err := stmt.QueryContext(Ctx, db, &messageFiles); err != nil {
		return nil, err
	}
	return messageFiles, nil
}
