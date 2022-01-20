package service

import "auth-service/model/request"

type Auth interface {
	Login(request request.User) string
}
