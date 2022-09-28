package controllers

import (
	"awesomeProject/services"
	u "awesomeProject/utils"
	"net/http"
)


var SyncStart = func(w http.ResponseWriter, r *http.Request) {
	resp := services.StartAll()
	u.Respond(w, resp)
}