package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	"errors"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"time"
)

const (
	ChatTypeSingle int64 = iota + 1
	ChatTypeGroup
)

const (
	ChatNameUniqueGeneral string = "1"
	ChatNameUniqueBranch  string = "2"
)

type Chat struct {
	model.Chats
	UnreadCount      int64
	ParticipantCount int64
	Messages         []model.Messages `alias:"ChatMessages.*"`
	LastMessage      *struct {
		model.Messages `alias:"ChatLastMessageLateral.*"`
		Sender         struct {
			model.Users `alias:"LastMessageSender.*"`
		}
	}
	Participants []model.Participants `alias:"chatParticipants.*"`
	Companion    struct {
		model.Users `alias:"chatCompanion.*"`
		LastOnline  *struct {
			model.Onlines `alias:"UserLastOnline.*"`
		}
	}
}

var ChatMessages = Messages.AS("ChatMessages")
var chatMessageSender = Users.AS("MessageSender")
var ChatLastMessage = Messages.AS("ChatLastMessageLateral")
var ChatLastMessageSender = Users.AS("LastMessageSender")
var chatParticipants = Participants.AS("chatParticipants")
var chatCompanion = Users.AS("chatCompanion")
var chatInsertedColumns = Chats.MutableColumns.Except(Chats.CreatedAt, Chats.DeletedAt, Chats.UpdatedAt)

func ChatInit(Type int64, CreatorID int64, Name *string) model.Chats {
	return model.Chats{
		Type:      Type,
		CreatorID: CreatorID,
		Name:      Name,
	}
}

func ChatFrom(userID int64) ReadableTable {
	ChatUnreadCountID := Messages.ChatID.From(ChatUnreadCount(userID))
	ChatParticipantCountID := Participants.ChatID.From(ChatParticipantCount())

	return Chats.LEFT_JOIN(ChatLastMessageLateral(), Bool(true)).
		INNER_JOIN(chatParticipants, chatParticipants.ChatID.EQ(Chats.ID)).
		LEFT_JOIN(ChatUnreadCount(userID), ChatUnreadCountID.EQ(Chats.ID)).
		INNER_JOIN(ChatParticipantCount(), ChatParticipantCountID.EQ(Chats.ID))
}

func ChatSelect(userID int64) ProjectionList {
	//companionAvatarId := Users.AvatarID.From(companion)
	//companionFullName := FloatColumn("FullName").From(companion)

	return ProjectionList{
		Chats.AllColumns, chatParticipants.AllColumns, ChatLastMessageLateral().AllColumns(),
		ChatUnreadCount(userID).AllColumns(), ChatParticipantCount().AllColumns(),
		//CASE().WHEN(Chats.Type.EQ(Int(1))).THEN(companionAvatarId).
		//	ELSE(NULL).AS("Chat.AvatarID"),
		//CASE().WHEN(Chats.Type.EQ(Int(1))).THEN(companionFullName).
		//	ELSE(Chats.Name).AS("Chat.name"),
	}
}

func ChatLastMessageLateral() SelectTable {
	return LATERAL(
		SELECT(
			ChatLastMessage.AllColumns, ChatLastMessageSender.AllColumns,
		).FROM(
			ChatLastMessage.
				INNER_JOIN(ChatLastMessageSender, ChatLastMessageSender.ID.EQ(ChatLastMessage.SenderID)),
		).WHERE(
			Chats.ID.EQ(ChatLastMessage.ChatID).AND(ChatLastMessage.DeletedAt.GT(RawTimestampz("now()"))),
		).ORDER_BY(ChatLastMessage.ID.DESC()).LIMIT(1),
	).AS("ChatLastMessageLateral")
}

func ChatLastReadMessageID(userID int64) SelectTable {

	return SELECT(
		Messages.ChatID,
		CASE().WHEN(MAX(Reads.MessageID).IS_NULL()).THEN(Int(0)).
			ELSE(MAX(Reads.MessageID)).AS("MessageID"),
	).FROM(
		Messages.LEFT_JOIN(Reads, Messages.ID.EQ(Reads.MessageID)),
	).WHERE(Reads.UserID.EQ(Int(userID))).
		GROUP_BY(Messages.ChatID).AsTable("ChatLastReadMessageID")
}

func ChatUnreadCount(userID int64) SelectTable {

	ChatLastReadChatID := Messages.ChatID.From(ChatLastReadMessageID(userID))
	chatLastReadMessageID := IntegerColumn("MessageID").From(ChatLastReadMessageID(userID))

	return SELECT(
		Messages.ChatID, COUNT(Messages.ChatID).AS("Chat.UnreadCount"),
	).FROM(
		Messages.LEFT_JOIN(ChatLastReadMessageID(userID), ChatLastReadChatID.EQ(Messages.ChatID)),
	).WHERE(Messages.ID.GT(chatLastReadMessageID).OR(chatLastReadMessageID.IS_NULL()).AND(Messages.SenderID.NOT_EQ(Int(userID)))).
		GROUP_BY(Messages.ChatID).AsTable("ChatUnreadCount")
}

func ChatParticipantCount() SelectTable {
	return SELECT(
		Participants.ChatID, COUNT(Participants.ChatID).AS("Chat.ParticipantCount"),
	).FROM(
		Participants,
	).GROUP_BY(Participants.ChatID).AsTable("ChatParticipantCount")
}

func ChatCompanion(userId int64) SelectTable {
	return LATERAL(
		SELECT(
			chatCompanion.AllColumns, chatCompanion.FirstName.CONCAT(String(" ")).CONCAT(chatCompanion.LastName).AS("FullName"),
			LastOnlineLateral(chatCompanion).AllColumns(),
		).FROM(
			Participants.
				INNER_JOIN(chatCompanion, chatCompanion.ID.EQ(Participants.UserID)).
				LEFT_JOIN(LastOnlineLateral(chatCompanion), Bool(true)),
		).WHERE(
			Chats.ID.EQ(Participants.ChatID).
				AND(Participants.UserID.NOT_EQ(Int(userId))),
		).LIMIT(1),
	).AS("ChatCompanion")
}

func GetChat(chatId int64, userId int64) (*Chat, error) {
	now := time.Now()
	iParticipant := Participants.AS("iParticipant")

	stmt := SELECT(
		ChatSelect(0),
	).FROM(ChatFrom(0).
		INNER_JOIN(iParticipant, iParticipant.ChatID.EQ(Chats.ID).
			AND(iParticipant.UserID.EQ(Int(userId))). // я получатель
			AND(iParticipant.DeletedAt.GT_EQ(TimestampzT(now)))),
	).WHERE(
		Chats.ID.EQ(Int(chatId)).AND(Chats.DeletedAt.GT_EQ(TimestampzT(now))),
	)

	var chats []Chat
	if err := stmt.Query(JetDB, &chats); err != nil {
		return nil, err
	}
	if chats == nil {
		return nil, nil
	}
	return &chats[0], nil
}

func InsertChat(db qrm.DB, chat model.Chats, participants []int64) (model.Chats, error) {
	if chat.Type == 1 {
		if len(participants) < 2 {
			return chat, errors.New("Для личного чата требуется 2 участника")
		}

		chat, err := GetSingleChat(participants[0], participants[1])
		if err == nil || err != qrm.ErrNoRows {
			return chat, nil
		}
	}

	stmt := Chats.INSERT(Chats.Type, Chats.CreatorID, Chats.Name).
		MODEL(chat).RETURNING(Chats.AllColumns)

	if err := stmt.QueryContext(Ctx, db, &chat); err != nil {
		return chat, err
	}

	_, err := InsertParticipants(db, chat.ID, participants)
	if err != nil {
		return chat, err
	}
	return chat, nil
}

func GetSingleChat(user1 int64, user2 int64) (model.Chats, error) {
	p1 := Participants.AS("p1")
	p2 := Participants.AS("p2")

	stmt := SELECT(
		Chats.AllColumns,
	).FROM(
		Chats.
			INNER_JOIN(p1, p1.ChatID.EQ(Chats.ID).AND(p1.UserID.EQ(Int(user1)))).
			INNER_JOIN(p2, p2.ChatID.EQ(Chats.ID).AND(p2.UserID.EQ(Int(user2)))),
	).WHERE(
		Chats.Type.EQ(Int(1)).AND(NotDeleted(Chats.DeletedAt)),
	).LIMIT(1)
	var chat model.Chats

	err := stmt.Query(JetDB, &chat)
	if err != nil {
		return chat, err
	}
	return chat, nil
}

func UpdateChat(db qrm.DB, chat *model.Chats) error {
	stmt := Chats.UPDATE(chatInsertedColumns).
		MODEL(chat).
		WHERE(Chats.ID.EQ(Int(chat.ID))).RETURNING(Chats.AllColumns)

	err := stmt.QueryContext(Ctx, db, chat)
	return err
}

func DeleteChats(db qrm.DB, ids []int64) ([]model.Chats, error) {
	stmt := Chats.UPDATE(Chats.DeletedAt).
		SET("now()").
		WHERE(Chats.ID.IN(InInt(ids)...)).
		RETURNING(Chats.AllColumns)

	var chats []model.Chats
	if err := stmt.QueryContext(Ctx, db, &chats); err != nil {
		return nil, err
	}
	return chats, nil
}
