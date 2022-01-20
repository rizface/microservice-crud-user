package controller

import (
	"auth-service/exception"
	"auth-service/helper"
	request2 "auth-service/model/request"
	"auth-service/service"
	"encoding/json"
	"net/http"
)

type auth struct {
	service service.Auth
}

func NewAuth(service service.Auth) Auth {
	return &auth{
		service: service,
	}
}

func (a *auth) Login(w http.ResponseWriter, r *http.Request) {
	var request request2.User
	err := json.NewDecoder(r.Body).Decode(&request)
	exception.InternalServerError(err)
	token := a.service.Login(request)
	helper.JsonWriter(w,http.StatusOK,"success",map[string]interface{} {
		"token": token,
	})
}

