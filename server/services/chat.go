package services

import (
	"awesomeProject/api"
	"awesomeProject/database"
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	"awesomeProject/internal"
	"awesomeProject/utils"
	. "github.com/ahmetb/go-linq/v3"
	. "github.com/go-jet/jet/v2/postgres"
	"time"
)

func GetMyChats(userId int64) map[string]interface{} {

	now := time.Now()
	iParticipant := Participants.AS("iParticipant")
	messageId := database.ChatLastMessage.ID.From(database.ChatLastMessageLateral())

	stmt := SELECT(
		database.ChatSelect(userId), database.ChatCompanion(userId).AllColumns(),
	).FROM(database.ChatFrom(userId).
		INNER_JOIN(iParticipant, iParticipant.ChatID.EQ(Chats.ID).
			AND(iParticipant.UserID.EQ(Int(userId))). // я получатель
			AND(iParticipant.DeletedAt.GT_EQ(TimestampzT(now)))).
		LEFT_JOIN(database.ChatCompanion(userId), Bool(true)),
	).WHERE(
		Chats.DeletedAt.GT_EQ(TimestampzT(now)).AND(NOT(Chats.Type.EQ(Int(1)).AND(messageId.IS_NULL()))),
	).ORDER_BY(messageId.DESC())
	//fmt.Println(stmt.DebugSql())
	var chats []database.Chat
	err := stmt.Query(database.JetDB, &chats)
	if err != nil {
		return internal.LogError(err.Error(), 824675)

	}

	resp := utils.Message(true, "")
	resp["Chats"] = chats
	return resp
}

func GetChat(getById api.ById) map[string]interface{} {
	now := time.Now()
	iParticipant := Participants.AS("iParticipant")
	messageId := database.ChatLastMessage.ID.From(database.ChatLastMessageLateral())

	stmt := SELECT(
		database.ChatSelect(getById.SenderId), database.ChatCompanion(getById.SenderId).AllColumns(),
	).FROM(database.ChatFrom(getById.SenderId).
		INNER_JOIN(iParticipant, iParticipant.ChatID.EQ(Chats.ID).
			AND(iParticipant.UserID.EQ(Int(getById.SenderId))). // я получатель
			AND(iParticipant.DeletedAt.GT_EQ(TimestampzT(now)))).
		LEFT_JOIN(database.ChatCompanion(getById.SenderId), Bool(true)),
	).WHERE(
		Chats.ID.EQ(Int(getById.Id)).AND(Chats.DeletedAt.GT_EQ(TimestampzT(now))),
	).ORDER_BY(messageId.DESC())

	var chat database.Chat
	err := stmt.Query(database.JetDB, &chat)
	if err != nil {
		return internal.LogError(err.Error(), 793678)
	}

	resp := utils.Message(true, "")
	resp["Chat"] = chat
	return resp
}

func ChatAddParticipants(participantsAdd api.ParticipantsAdd) map[string]interface{} {

	stmt := SELECT(
		database.ChatSelect(0),
	).FROM(
		database.ChatFrom(0),
	).WHERE(
		Chats.ID.EQ(Int(participantsAdd.ChatId)).
			AND(Chats.CreatorID.EQ(Int(participantsAdd.SenderId))).
			AND(Chats.DeletedAt.GT(RawTimestampz("now()"))),
	)

	var chat database.Chat
	err := stmt.Query(database.JetDB, &chat)
	if err != nil {
		return internal.LogError(err.Error(), 824965)

	}

	_, err = database.InsertParticipants(database.JetDB, participantsAdd.ChatId, participantsAdd.Ids)
	if err != nil {
		return internal.LogError(err.Error(), 244839)
	}

	internal.SocketSendUpdate(database.ParticipantsGetUserIDs(chat.Participants))

	resp := utils.Message(true, "")
	return resp
}

func ChatDelParticipants(participantsAdd api.ParticipantsAdd) map[string]interface{} {

	participants, err := database.DeleteParticipants(database.JetDB, participantsAdd.ChatId, participantsAdd.Ids)
	if err != nil {
		return internal.LogError(err.Error(), 924831)
	}

	resp := utils.Message(true, "")
	resp["Participants"] = participants
	return resp
}

func GetMessages(getMessages api.GetMessages) map[string]interface{} {
	//count := pointy.Int64Value(getMessages.Count, 10)
	//startMessageId := pointy.Int64Value(getMessages.StartMessageId, 999999999)

	stmt := SELECT(
		database.MessageSelect(getMessages.SenderId),
		//CASE().WHEN(message.SenderID.EQ(Int(userId))).THEN(Bool(true)).
		//	ELSE(Bool(false)).AS("Message.IsMy"),
	).FROM(
		database.MessageFrom(getMessages.SenderId),
	).WHERE(
		Messages.ChatID.EQ(Int(getMessages.ChatId)).
			AND(database.NotDeleted(Messages.DeletedAt)),
	).ORDER_BY(Messages.ID.DESC())
	var messages []database.Message
	//println(stmt.DebugSql())
	if err := stmt.Query(database.JetDB, &messages); err != nil {
		return internal.LogError(err.Error(), 917062)
	}

	resp := utils.Message(true, "")
	resp["Messages"] = messages
	return resp
}

func SetActivity(activity model.Activitys) map[string]interface{} {

	now := time.Now()
	iParticipant := Participants.AS("iParticipant")

	stmt := SELECT(
		database.ChatSelect(0),
	).FROM(database.ChatFrom(0).
		INNER_JOIN(iParticipant, iParticipant.ChatID.EQ(Chats.ID).
			AND(iParticipant.UserID.EQ(Int(activity.UserID))). // я получатель
			AND(iParticipant.DeletedAt.GT_EQ(TimestampzT(now)))),
	).WHERE(
		Chats.ID.EQ(Int(activity.ChatID)).AND(Chats.DeletedAt.GT_EQ(TimestampzT(now))),
	)

	var chat database.Chat
	err := stmt.Query(database.JetDB, &chat)
	if err != nil {
		return internal.LogError(err.Error(), 105927)
	}

	activity, err = database.InsertActivity(database.JetDB, activity)
	if err != nil {
		return internal.LogError(err.Error(), 523057)
	}

	resp := utils.Message(true, "")
	resp["Activity"] = activity
	return resp
}

func ChatEdit(edit api.ChatEdit) map[string]interface{} {

	var chat database.Chat
	err := SELECT(
		database.ChatSelect(edit.SenderId),
	).FROM(
		database.ChatFrom(edit.SenderId),
	).WHERE(
		Chats.ID.EQ(Int(edit.ChatID)),
	).Query(database.JetDB, &chat)
	if err != nil {
		return internal.LogError(err.Error(), 208667)
	}

	if edit.Name != nil {
		chat.Name = utils.StringEmptyToNil(edit.Name)
	}
	if edit.About != nil {
		chat.About = utils.StringEmptyToNil(edit.About)
	}
	if edit.AvatarID != nil {
		chat.AvatarID = utils.IntEmptyToNil(edit.AvatarID)
	}

	err = database.UpdateChat(database.JetDB, &chat.Chats)
	if err != nil {
		return internal.LogError(err.Error(), 814279)
	}

	internal.SocketSendUpdate(database.ParticipantsGetUserIDs(chat.Participants))

	resp := utils.Message(true, "")
	resp["Chat"] = chat
	return resp
}

func ChatsDelete(byIds api.Request) map[string]interface{} {
	iParticipant := Participants.AS("iParticipant")

	stmt := SELECT(
		database.ChatSelect(byIds.SenderId),
	).FROM(
		database.ChatFrom(byIds.SenderId).
			INNER_JOIN(iParticipant, iParticipant.ChatID.EQ(Chats.ID).
				AND(iParticipant.UserID.EQ(Int(byIds.SenderId))). // я получатель
				AND(iParticipant.DeletedAt.GT_EQ(RawTimestampz("now()")))),
	).WHERE(
		Chats.ID.IN(database.InInt(byIds.Ids)...).
			AND(Chats.Type.EQ(Int(database.ChatTypeSingle))).
			AND(Chats.DeletedAt.GT(RawTimestampz("now()"))),
	)

	var chats []database.Chat
	err := stmt.Query(database.JetDB, &chats)
	if err != nil {
		return internal.LogError(err.Error(), 278058)
	}

	var chatIDs []int64
	From(chats).Select(func(c interface{}) interface{} {
		return c.(database.Chat).ID
	}).ToSlice(&chatIDs)

	_, err = database.DeleteChats(database.JetDB, chatIDs)
	if err != nil {
		return internal.LogError(err.Error(), 140658)
	}

	internal.SocketSendUpdate([]int64{byIds.SenderId})

	resp := utils.Message(true, "")
	return resp
}

func CreateChat(createChat api.CreateChat) map[string]interface{} {
	participantIds := utils.RemoveDuplicateValues(append(createChat.Participants, createChat.SenderId))
	chat := database.ChatInit(createChat.Type, createChat.SenderId, createChat.Name)

	tx, _ := database.JetDB.Begin()
	chat, err := database.InsertChat(tx, chat, participantIds)
	if err != nil {
		_ = tx.Rollback()
		return internal.LogError(err.Error(), 805623)
	}
	_ = tx.Commit()

	if chat.Type == database.ChatTypeGroup {
		internal.SocketSendUpdate(participantIds)
	}

	res := utils.Message(true, "")
	res["Chat"] = chat

	return res
}
