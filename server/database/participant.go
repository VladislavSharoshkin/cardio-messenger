package database

import (
	"awesomeProject/gen/postgres/public/model"
	. "awesomeProject/gen/postgres/public/table"
	. "github.com/go-jet/jet/v2/postgres"

	"github.com/go-jet/jet/v2/qrm"
)

type Participant struct {
	model.Participants
	User struct {
		model.Users `alias:"ParticipantUser.*"`
	}
}

var participantUser = Users.AS("ParticipantUser")

func ParticipantsGetUserIDs(participants []model.Participants) []int64 {
	var userIDs []int64
	for _, par := range participants {
		userIDs = append(userIDs, par.UserID)
	}
	return userIDs
}

func ParticipantSelect() ProjectionList {
	return ProjectionList{
		Participants.AllColumns, participantUser.AllColumns,
	}
}

func ParticipantFrom() ReadableTable {
	return Participants.INNER_JOIN(participantUser, participantUser.ID.EQ(Participants.UserID))
}

func InsertParticipants(db qrm.DB, chatId int64, userIds []int64) ([]model.Participants, error) {
	var participants []model.Participants
	for _, userId := range userIds {
		participants = append(participants, model.Participants{ChatID: chatId, UserID: userId})
	}

	stmt := Participants.INSERT(Participants.UserID, Participants.ChatID).
		MODELS(&participants).RETURNING(Participants.AllColumns)

	if err := stmt.QueryContext(Ctx, db, &participants); err != nil {
		return nil, err
	}
	return participants, nil
}

func DeleteParticipants(db qrm.DB, chatID int64, ids []int64) ([]model.Participants, error) {
	stmt := Participants.UPDATE(Participants.DeletedAt).
		SET("now()").
		WHERE(Participants.ChatID.EQ(Int(chatID)).AND(Participants.UserID.IN(InInt(ids)...)))

	var participants []model.Participants
	if err := stmt.QueryContext(Ctx, db, &participants); err != nil {
		return nil, err
	}
	return participants, nil
}
