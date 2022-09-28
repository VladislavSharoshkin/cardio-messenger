package controllers

import (
	"awesomeProject/api"
	"awesomeProject/services"
	u "awesomeProject/utils"
	"encoding/json"
	"net/http"
)

// var AddRight = func(w http.ResponseWriter, r *http.Request) {
//
//		right := models.Right{}
//		err := json.NewDecoder(r.Body).Decode(&right)
//		if err != nil {
//			u.Respond(w, u.ErrorInit(u.ErrorCodeDecodeRequest, err.Error(), 551264).ToResponse())
//			return
//		}
//
//		resp := models.AddRights([]models.Right{right})
//		u.Respond(w, resp)
//	}
//
// var AddRoleType = func(w http.ResponseWriter, r *http.Request) {
//
//		roleType := models.RoleType{}
//		err := json.NewDecoder(r.Body).Decode(&roleType)
//		if err != nil {
//			u.Respond(w, u.ErrorInit(u.ErrorCodeDecodeRequest, err.Error(), 364633).ToResponse())
//			return
//		}
//
//		resp := models.AddRoleTypes([]models.RoleType{roleType})
//		u.Respond(w, resp)
//	}
//
// var AddRole = func(w http.ResponseWriter, r *http.Request) {
//
//		role := models.Role{}
//		err := json.NewDecoder(r.Body).Decode(&role)
//		if err != nil {
//			u.Respond(w, u.ErrorInit(u.ErrorCodeDecodeRequest, err.Error(), 582457).ToResponse())
//			return
//		}
//
//		resp := models.AddRoles([]models.Role{role})
//		u.Respond(w, resp)
//	}

var UserAdd = func(w http.ResponseWriter, r *http.Request) {

	reg := api.Reg{}
	err := json.NewDecoder(r.Body).Decode(&reg)
	if err != nil {
		u.Respond(w, u.ErrorInit(u.ErrorCodeDecodeRequest, err.Error(), 713981).ToResponse())
		return
	}

	resp := services.UserAdd(reg)
	u.Respond(w, resp)
}
