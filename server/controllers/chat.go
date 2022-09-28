package controllers

import (
	"awesomeProject/api"
	"awesomeProject/gen/postgres/public/model"
	"awesomeProject/internal"
	"awesomeProject/services"
	u "awesomeProject/utils"
	"encoding/json"
	"net/http"
)

var ChatCreate = func(w http.ResponseWriter, r *http.Request) {

	createChat := api.CreateChat{}
	err := json.NewDecoder(r.Body).Decode(&createChat)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 925935))
		return
	}

	createChat.SenderId = r.Context().Value("user").(int64)

	resp := services.CreateChat(createChat)
	u.Respond(w, resp)
}

var GetMyChats = func(w http.ResponseWriter, r *http.Request) {

	var userId = r.Context().Value("user").(int64)

	resp := services.GetMyChats(userId)
	u.Respond(w, resp)
}

var GetChat = func(w http.ResponseWriter, r *http.Request) {

	getById := api.ById{}
	err := json.NewDecoder(r.Body).Decode(&getById)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 393617))
		return
	}
	getById.SenderId = r.Context().Value("user").(int64)

	resp := services.GetChat(getById)
	u.Respond(w, resp)
}



var EditChat = func(w http.ResponseWriter, r *http.Request) {

	edit := api.ChatEdit{}
	err := json.NewDecoder(r.Body).Decode(&edit)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 719568))
		return
	}

	edit.SenderId = r.Context().Value("user").(int64)

	resp := services.ChatEdit(edit)
	u.Respond(w, resp)
}

var DeleteChats = func(w http.ResponseWriter, r *http.Request) {

	getByIds := api.Request{}
	err := json.NewDecoder(r.Body).Decode(&getByIds)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 287046))
		return
	}

	getByIds.SenderId = r.Context().Value("user").(int64)

	resp := services.ChatsDelete(getByIds)
	u.Respond(w, resp)
}

var ParticipantAdd = func(w http.ResponseWriter, r *http.Request) {

	participantsAdd := api.ParticipantsAdd{}
	err := json.NewDecoder(r.Body).Decode(&participantsAdd)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 834803))
		return
	}

	participantsAdd.SenderId = r.Context().Value("user").(int64)

	resp := services.ChatAddParticipants(participantsAdd)
	u.Respond(w, resp)
}

var ParticipantDel = func(w http.ResponseWriter, r *http.Request) {

	participantsAdd := api.ParticipantsAdd{}
	err := json.NewDecoder(r.Body).Decode(&participantsAdd)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 357892))
		return
	}

	participantsAdd.SenderId = r.Context().Value("user").(int64)

	resp := services.ChatDelParticipants(participantsAdd)
	u.Respond(w, resp)
}

var GetParticipants = func(w http.ResponseWriter, r *http.Request) {

	byId := api.ById{}
	err := json.NewDecoder(r.Body).Decode(&byId)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 657025))
		return
	}

	byId.SenderId = r.Context().Value("user").(int64)

	resp := services.GetParticipants(byId)
	u.Respond(w, resp)
}

var SetActivity = func(w http.ResponseWriter, r *http.Request) {

	activity := model.Activitys{}
	err := json.NewDecoder(r.Body).Decode(&activity)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 160396))
		return
	}

	activity.UserID = r.Context().Value("user").(int64)

	resp := services.SetActivity(activity)
	u.Respond(w, resp)
}
