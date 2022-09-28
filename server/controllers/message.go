package controllers

import (
	"awesomeProject/api"
	"awesomeProject/internal"
	"awesomeProject/services"
	u "awesomeProject/utils"
	"encoding/json"
	"net/http"
)

var Send = func(w http.ResponseWriter, r *http.Request) {

	send := api.SendMessage{}
	err := json.NewDecoder(r.Body).Decode(&send)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 424505))
		return
	}
	send.SenderId = r.Context().Value("user").(int64)

	resp := services.MessageSend(send)
	u.Respond(w, resp)
}

var GetMessages = func(w http.ResponseWriter, r *http.Request) {

	getMessages := api.GetMessages{}
	err := json.NewDecoder(r.Body).Decode(&getMessages)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 825486))
		return
	}
	getMessages.SenderId = r.Context().Value("user").(int64)

	resp := services.GetMessages(getMessages)
	u.Respond(w, resp)
}


var DeleteMessagesByIds = func(w http.ResponseWriter, r *http.Request) {

	deleteMessage := api.DeleteMessage{}
	err := json.NewDecoder(r.Body).Decode(&deleteMessage)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 925935))
		return
	}

	deleteMessage.SenderId = r.Context().Value("user").(int64)

	resp := services.MessagesDelete(deleteMessage)
	u.Respond(w, resp)
}

var EditMessage = func(w http.ResponseWriter, r *http.Request) {

	edit := api.MessageEdit{}
	err := json.NewDecoder(r.Body).Decode(&edit)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 108581))
		return
	}

	edit.SenderId = r.Context().Value("user").(int64)

	resp := services.MessageEdit(edit)
	u.Respond(w, resp)
}

var ReadMessage = func(w http.ResponseWriter, r *http.Request) {

	readMessage := api.ReadMessage{}
	err := json.NewDecoder(r.Body).Decode(&readMessage)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 924394))
		return
	}

	readMessage.SenderId = r.Context().Value("user").(int64)

	resp := services.MessageRead(readMessage)
	u.Respond(w, resp)
}