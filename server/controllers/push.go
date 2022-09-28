package controllers

import (
	"awesomeProject/api"
	"awesomeProject/internal"
	"awesomeProject/services"
	u "awesomeProject/utils"
	"encoding/json"
	"net/http"
)

var PushAdd = func(w http.ResponseWriter, r *http.Request) {

	pushAdd := api.PushAdd{}
	err := json.NewDecoder(r.Body).Decode(&pushAdd)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 379359))
		return
	}

	resp := services.PushAdd(pushAdd)
	u.Respond(w, resp)
}

var PushSend = func(w http.ResponseWriter, r *http.Request) {

	pushAdd := api.PushSend{}
	err := json.NewDecoder(r.Body).Decode(&pushAdd)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 626828))
		return
	}

	resp := services.PushSend(pushAdd)
	u.Respond(w, resp)
}