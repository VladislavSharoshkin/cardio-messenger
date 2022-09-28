package services

import (
	"awesomeProject/api"
	"awesomeProject/database"
	. "awesomeProject/gen/postgres/public/table"
	"awesomeProject/internal"
	"awesomeProject/utils"
	"github.com/ahmetb/go-linq/v3"
	. "github.com/go-jet/jet/v2/postgres"
	"time"
)

func MessageSend(sendMessage api.SendMessage) map[string]interface{} {
	if ok, res := sendMessage.Validate(); !ok {
		return res
	}

	var chat database.Chat
	iParticipant := Participants.AS("iParticipant")
	err := SELECT(
		database.ChatSelect(0),
	).FROM(
		database.ChatFrom(0).INNER_JOIN(iParticipant, iParticipant.ChatID.EQ(Chats.ID).
			AND(iParticipant.UserID.EQ(Int(sendMessage.SenderId))).
			AND(iParticipant.DeletedAt.GT_EQ(TimestampzT(time.Now())))),
	).WHERE(
		Chats.ID.EQ(Int(sendMessage.ChatId)).
			AND(database.NotDeleted(Chats.DeletedAt)),
	).Query(database.JetDB, &chat)
	if err != nil {
		return internal.LogError(err.Error(), 931959)
	}
	message := database.MessageInit(chat.ID, sendMessage.SenderId, sendMessage.Text)
	tx, _ := database.JetDB.Begin()
	err = database.InsertMessage(tx, &message, sendMessage.ForwardMessages, sendMessage.Attachments)
	if err != nil {
		tx.Rollback()
		return internal.LogError(err.Error(), 544298)
	}
	tx.Commit()

	participantUserIds := database.ParticipantsGetUserIDs(chat.Participants)

	internal.SocketSendUpdate(participantUserIds)

	sender, err := database.UserById(sendMessage.SenderId)
	if err != nil {
		return internal.LogError(err.Error(), 427979)
	}

	var participantWithoutSender []int64

	linq.From(participantUserIds).
		Except(linq.From([]int64{sendMessage.SenderId})).
		ToSlice(&participantWithoutSender)

	PushSend2(participantWithoutSender, sender.FirstName+" "+sender.LastName, utils.StringValue(sendMessage.Text))

	response := utils.Message(true, "")
	response["Message"] = message
	return response
}

func MessageRead(messageRead api.ReadMessage) map[string]interface{} {
	iParticipant := Participants.AS("iParticipant")

	stmt := SELECT(
		database.ChatSelect(messageRead.SenderId),
	).FROM(
		database.ChatFrom(messageRead.SenderId).INNER_JOIN(iParticipant, iParticipant.ChatID.EQ(Chats.ID).
			AND(iParticipant.UserID.EQ(Int(messageRead.SenderId))).
			AND(iParticipant.DeletedAt.GT_EQ(TimestampzT(time.Now())))).
			INNER_JOIN(Messages, Messages.ID.EQ(Int(messageRead.MessageId)).
				AND(Messages.DeletedAt.GT(RawTimestampz("now()")))),
	).WHERE(
		Chats.ID.EQ(Int(messageRead.ChatId)).AND(database.NotDeleted(Chats.DeletedAt)),
	)
	var chat database.Chat
	err := stmt.Query(database.JetDB, &chat)

	read := database.ReadInit(messageRead.MessageId, messageRead.SenderId)
	err = database.InsertRead(database.JetDB, &read)
	if err != nil {
		return internal.LogError(err.Error(), 145844)
	}

	internal.SocketSendUpdate(database.ParticipantsGetUserIDs(chat.Participants))

	response := utils.Message(true, "")
	response["Read"] = read
	return response
}

func MessagesDelete(byIds api.DeleteMessage) map[string]interface{} {

	stmt := SELECT(
		database.ChatSelect(byIds.SenderId), database.ChatMessages.AllColumns,
	).FROM(
		database.ChatFrom(byIds.SenderId).
			LEFT_JOIN(database.ChatMessages, database.ChatMessages.ChatID.EQ(Chats.ID)),
	).WHERE(
		Chats.ID.EQ(Int(byIds.ChatId)).
			AND(Chats.DeletedAt.GT(RawTimestampz("now()"))).
			AND(database.ChatMessages.ID.IN(database.InInt(byIds.Ids)...)).
			AND(database.ChatMessages.SenderID.EQ(Int(byIds.SenderId))),
	)

	var chat database.Chat
	err := stmt.Query(database.JetDB, &chat)
	if err != nil {
		return internal.LogError(err.Error(), 724674)
	}

	messages, err := database.DeleteMessages(database.JetDB, chat.ID, database.MessagesGetIDs(chat.Messages))
	if err != nil {
		return internal.LogError(err.Error(), 965274)
	}

	internal.SocketSendUpdate(database.ParticipantsGetUserIDs(chat.Participants))

	resp := utils.Message(true, "")
	resp["Messages"] = messages
	return resp
}

func MessageEdit(messageEdit api.MessageEdit) map[string]interface{} {

	var message database.Message
	err := SELECT(
		database.MessageSelect(0),
	).FROM(
		database.MessageFrom(0),
	).WHERE(
		Messages.ID.EQ(Int(messageEdit.ID)).
			AND(Messages.SenderID.EQ(Int(messageEdit.SenderId))).
			AND(Messages.DeletedAt.GT(RawTimestampz("now()"))),
	).Query(database.JetDB, &message)
	if err != nil {
		return internal.LogError(err.Error(), 258094)
	}

	if messageEdit.Text != nil {
		message.Text = utils.StringEmptyToNil(messageEdit.Text)
	}

	err = database.UpdateMessage(database.JetDB, &message.Messages)
	if err != nil {
		return internal.LogError(err.Error(), 611495)
	}

	resp := utils.Message(true, "")
	resp["Message"] = message
	return resp
}
