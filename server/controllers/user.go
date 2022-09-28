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

var Reg = func(w http.ResponseWriter, r *http.Request) {
	reg := api.Reg{}
	err := json.NewDecoder(r.Body).Decode(&reg) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 362193))
		return
	}

	resp := services.Reg(reg) //Создать аккаунт
	u.Respond(w, resp)
}

var Search = func(w http.ResponseWriter, r *http.Request) {

	searchUser := api.SearchUsers{}
	err := json.NewDecoder(r.Body).Decode(&searchUser)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 393901))
		return
	}

	resp := services.SearchUsers(searchUser)
	u.Respond(w, resp)
}

var Login = func(w http.ResponseWriter, r *http.Request) {

	login := api.Login{}
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 319738))
		return
	}

	resp := services.Login(login)
	u.Respond(w, resp)
}

var Logout = func(w http.ResponseWriter, r *http.Request) {

	byId := api.ById{}
	err := json.NewDecoder(r.Body).Decode(&byId)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 296693))
		return
	}

	byId.SenderId = r.Context().Value("user").(int64)

	resp := services.Logout(byId)
	u.Respond(w, resp)
}

var UserGet = func(w http.ResponseWriter, r *http.Request) {

	getById := api.ById{}
	err := json.NewDecoder(r.Body).Decode(&getById)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 225482))
		return
	}

	resp := services.GetUser(getById)
	u.Respond(w, resp)
}

var EditUser = func(w http.ResponseWriter, r *http.Request) {

	userEdit := api.UserEdit{}
	err := json.NewDecoder(r.Body).Decode(&userEdit)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 917936))
		return
	}

	userEdit.SenderId = r.Context().Value("user").(int64)

	resp := services.UserEdit(userEdit)
	u.Respond(w, resp)
}



var SetOnline = func(w http.ResponseWriter, r *http.Request) {

	online := model.Onlines{}
	err := json.NewDecoder(r.Body).Decode(&online)
	if err != nil {
		u.Respond(w, internal.LogError(err.Error(), 914847))
		return
	}

	online.UserID = r.Context().Value("user").(int64)

	resp := services.SetOnline(online)
	u.Respond(w, resp)
}