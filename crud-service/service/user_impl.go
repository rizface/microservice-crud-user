package service

import (
	"context"
	"crud-service/exception"
	"crud-service/helper"
	"crud-service/model/request"
	"crud-service/model/response"
	"crud-service/repository"
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

type user struct {
	db *gorm.DB
	valid *validator.Validate
	repo repository.User
}

func NewUser(db *gorm.DB, valid *validator.Validate, repo repository.User) User {
	return &user{
		db:    db,
		valid: valid,
		repo:  repo,
	}
}

func (u *user) Get() *[]response.User {
	ctx,cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	users,err := u.repo.Get(ctx,u.db)
	exception.InternalServerError(err)
	return users
}

func (u *user) GetById(userId string) *response.User {
	ctx,cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	user,err := u.repo.GetById(ctx,u.db,userId)
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound) {
			exception.NotFound(err,"data tidak ditemukan")
		} else {
			exception.InternalServerError(err)
		}
	}
	return user
}

func (u *user) Post(request request.User) *response.User {
	err := u.valid.Struct(request)
	exception.Validate(err)

	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	request.Password = helper.GeneratePassword(request.Password)
	result,err := u.repo.Post(ctx,u.db,request)
	exception.InternalServerError(err)
	return result
}

func (u *user) Delete(userId string) string {
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	success,err := u.repo.Delete(ctx,u.db,userId)
	exception.InternalServerError(err)
	if success == false{
		exception.NotFound(errors.New("data tidak ditemukan"),"data tidak ditemukan")
	}
	return "data berhasil dihapus"
}

func (u *user) Update(request request.User) string {
	err := u.valid.Struct(request)
	exception.Validate(err)

	if len(request.Password) == 0 {
		user := u.GetById(request.Id)
		request.Password = user.Password
	} else {
		request.Password = helper.GeneratePassword(request.Password)
	}

	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	success,err := u.repo.Update(ctx,u.db,request)
	exception.InternalServerError(err)

	if success {
		return "data berhasil diupdate"
	}
	return "data gagal diupdate"
}



