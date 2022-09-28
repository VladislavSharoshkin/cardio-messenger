package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type Message struct {
	model.Messages
	ReadState int64
	Sender    struct {
		model.Users `alias:"MessageSender.*"`
	}
	Chat struct {
		model.Chats `alias:"MessageChat.*"`
	}
	ForwardMessages []struct {
		model.ForwardMessages `alias:"MessageForwardMessages.*"`
		ForwardMessage        struct {
			model.Messages `alias:"ForwardMessage.*"`
		}
	}
	Attachments []struct {
		model.MessageFiles `alias:"MessageAttachments.*"`
		File               struct {
			model.Files `alias:"AttachmentsFile.*"`
		}
	}
}

var messageSender = Users.AS("MessageSender")
var messageForwardMessages = ForwardMessages.AS("MessageForwardMessages")
var forwardMessage = Messages.AS("ForwardMessage")
var messageAttachments = MessageFiles.AS("MessageAttachments")
var attachmentFile = Files.AS("AttachmentsFile")
var MessageChat = Chats.AS("MessageChat")
var messageInsertedColumns = Messages.MutableColumns.Except(Messages.CreatedAt, Messages.DeletedAt, Messages.UpdatedAt)

func MessageInit(ChatID int64, SenderID int64, Text *string) model.Messages {
	return model.Messages{
		ChatID:   ChatID,
		SenderID: SenderID,
		Text:     Text,
	}
}

func MessagesGetIDs(messages []model.Messages) []int64 {
	var mesIDs []int64
	for _, mes := range messages {
		mesIDs = append(mesIDs, mes.ID)
	}
	return mesIDs
}

func MessageSelect(userID int64) ProjectionList {
	messageLastReadMessageID := IntegerColumn("MessageID").From(MessageLastReadMessageID(userID))

	return ProjectionList{
		CASE().WHEN(messageLastReadMessageID.IS_NULL().OR(Messages.ID.GT(messageLastReadMessageID))).THEN(Int(1)).
			ELSE(Int(2)).AS("Message.ReadState"),
		Messages.AllColumns, messageSender.AllColumns, messageForwardMessages.AllColumns,
		forwardMessage.AllColumns, messageAttachments.AllColumns, attachmentFile.AllColumns,
		MessageLastReadMessageID(userID).AllColumns(), MessageChat.AllColumns,
	}
}

func MessageFrom(userID int64) ReadableTable {
	messageLastReadChatID := Messages.ChatID.From(MessageLastReadMessageID(userID))

	return Messages.INNER_JOIN(messageSender, messageSender.ID.EQ(Messages.SenderID)).
		LEFT_JOIN(messageForwardMessages, messageForwardMessages.MessageID.EQ(Messages.ID)).
		LEFT_JOIN(forwardMessage, forwardMessage.ID.EQ(messageForwardMessages.ForwardMessageID)).
		LEFT_JOIN(messageAttachments, messageAttachments.MessageID.EQ(Messages.ID)).
		LEFT_JOIN(attachmentFile, attachmentFile.ID.EQ(messageAttachments.FileID)).
		LEFT_JOIN(MessageLastReadMessageID(userID), messageLastReadChatID.EQ(Messages.ChatID)).
		INNER_JOIN(MessageChat, MessageChat.ID.EQ(Messages.ChatID))
}

func MessageLastReadMessageID(userID int64) SelectTable {
	//userID := 2458

	return SELECT(
		Messages.ChatID,
		CASE().WHEN(MAX(Reads.MessageID).IS_NULL()).THEN(Int(0)).
			ELSE(MAX(Reads.MessageID)).AS("MessageID"),
	).FROM(
		Messages.LEFT_JOIN(Reads, Messages.ID.EQ(Reads.MessageID)),
	).WHERE(Reads.UserID.NOT_EQ(Int(userID))).
		GROUP_BY(Messages.ChatID).AsTable("ChatLastReadMessageID")
}

func InsertMessage(db qrm.DB, message *model.Messages, forwardIds []int64, fileIds []int64) error {
	stmt := Messages.INSERT(messageInsertedColumns).
		MODEL(message).RETURNING(Messages.AllColumns)

	if err := stmt.QueryContext(Ctx, db, message); err != nil {
		return err
	}

	if len(forwardIds) != 0 {
		_, err := InsertForwardMessages(db, message.ID, forwardIds)
		if err != nil {
			return err
		}
	}

	if len(fileIds) != 0 {
		_, err := InsertMessageFiles(db, message.ID, fileIds)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteMessages(db qrm.DB, chatID int64, ids []int64) ([]model.Messages, error) {
	stmt := Messages.UPDATE(Messages.DeletedAt).
		SET("now()").
		WHERE(Messages.ChatID.EQ(Int(chatID)).
			AND(Messages.ID.IN(InInt(ids)...))).
		RETURNING(Messages.AllColumns)

	var messages []model.Messages
	if err := stmt.QueryContext(Ctx, db, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

func UpdateMessage(db qrm.DB, message *model.Messages) error {
	stmt := Messages.UPDATE(messageInsertedColumns).
		MODEL(message).
		WHERE(Messages.ID.EQ(Int(message.ID))).RETURNING(Messages.AllColumns)

	err := stmt.QueryContext(Ctx, db, message)
	return err
}
