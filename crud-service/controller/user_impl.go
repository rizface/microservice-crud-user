package controller

import (
	"crud-service/exception"
	"crud-service/helper"
	request2 "crud-service/model/request"
	"crud-service/service"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type user struct {
	service service.User
}

func NewUser(service service.User) User{
	return &user{service: service}
}

func (u *user) Get(w http.ResponseWriter, r *http.Request) {
	users := u.service.Get()
	helper.JsonWriter(w,http.StatusOK,"success",map[string]interface{} {
		"users":users,
	})
}

func (u *user) GetById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r,"userId")
	user := u.service.GetById(userId)
	helper.JsonWriter(w,http.StatusOK,"success",map[string]interface{} {
		"user":user,
	})
}

func (u *user) Post(w http.ResponseWriter, r *http.Request) {
	var request request2.User
	err := json.NewDecoder(r.Body).Decode(&request)
	exception.InternalServerError(err)
	user := u.service.Post(request)
	helper.JsonWriter(w,http.StatusOK,"success",map[string]interface{} {
		"user":user,
	})
}

func (u *user) Update(w http.ResponseWriter, r *http.Request) {
	var request request2.User
	err := json.NewDecoder(r.Body).Decode(&request)
	exception.InternalServerError(err)
	userId := chi.URLParam(r,"userId")
	request.Id = userId
	result := u.service.Update(request)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

func (u *user) Delete(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r,"userId")
	result := u.service.Delete(userId)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

func (u *user) Profile(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)
	user := *u.service.GetById(userId)
	helper.JsonWriter(w,http.StatusOK,"success",map[string]interface{} {
		"user":user,
	})
}

