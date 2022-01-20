package service

import (
	"auth-service/exception"
	"auth-service/helper"
	"auth-service/model/request"
	"auth-service/repository"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type auth struct {
	db *gorm.DB
	repo repository.User
}

func NewAuth(db *gorm.DB, repo repository.User) Auth {
	return &auth{
		db:   db,
		repo: repo,
	}
}

func (a *auth) Login(request request.User) string {
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	user,err := a.repo.FindByUsername(ctx,a.db,request.Username)
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound) {
			exception.NotFound(err,"akun tidak terdaftar")
		} else {
			exception.InternalServerError(err)
		}
	}
	err = helper.ComparePassword(request.Password,user.Password)
	exception.Unauthorized(err,"username / password kamu salah")
	token := helper.GenerateToken(*user)
	return token
}

