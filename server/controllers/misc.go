package controllers

import (
	"awesomeProject/services"
	u "awesomeProject/utils"
	"net/http"
)

var Info = func(w http.ResponseWriter, r *http.Request) {
	resp := services.Info() //Создать аккаунт
	u.Respond(w, resp)
}