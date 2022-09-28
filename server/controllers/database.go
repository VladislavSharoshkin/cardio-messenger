package controllers

import (
	"awesomeProject/api"
	"awesomeProject/internal"
	"awesomeProject/services"
	u "awesomeProject/utils"
	"encoding/json"
	"net/http"
)

var GetBranches = func(w http.ResponseWriter, r *http.Request) {

	getMessages := api.GetMessages{}
	err := json.NewDecoder(r.Body).Decode(&getMessages)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 825486))
		return
	}
	getMessages.SenderId = r.Context().Value("user").(int64)

	resp := services.GetBranches()
	u.Respond(w, resp)
}

var GetStaffs = func(w http.ResponseWriter, r *http.Request) {

	getMessages := api.GetMessages{}
	err := json.NewDecoder(r.Body).Decode(&getMessages)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 825762))
		return
	}
	getMessages.SenderId = r.Context().Value("user").(int64)

	resp := services.GetStaffs()
	u.Respond(w, resp)
}