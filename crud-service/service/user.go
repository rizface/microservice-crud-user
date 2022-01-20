package service

import (
	"crud-service/model/request"
	"crud-service/model/response"
)

type User interface {
	Get() *[]response.User
	GetById(userId string) *response.User
	Post(request request.User) *response.User
	Delete(userId string) string
	Update(request request.User) string
}