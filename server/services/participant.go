package services

import (
	"awesomeProject/api"
	"awesomeProject/internal"
	"time"

	"awesomeProject/database"
	. "awesomeProject/gen/postgres/public/table"
	. "github.com/go-jet/jet/v2/postgres"

	"awesomeProject/utils"
)

func GetParticipants(byId api.ById) map[string]interface{} {

	stmt := SELECT(database.ParticipantSelect()).
		FROM(database.ParticipantFrom()).
		WHERE(
			Participants.ChatID.EQ(Int(byId.Id)).
				AND(Participants.DeletedAt.GT_EQ(TimestampzT(time.Now()))))

	var participants []database.Participant
	err := stmt.Query(database.JetDB, &participants)
	if err != nil {
		return internal.LogError(err.Error(), 368946)
	}

	resp := utils.Message(true, "")
	resp["Participants"] = participants
	return resp
}
