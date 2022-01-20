package repository

import (
	"context"
	"crud-service/model/app"
	"crud-service/model/request"
	"crud-service/model/response"
	"gorm.io/gorm"
)

type user struct{}

func NewUser() User {
	return &user{}
}

func (u *user) Get(ctx context.Context, db *gorm.DB) (*[]response.User,error) {
	var users []response.User
	result := db.WithContext(ctx).Select("id,username,level").Find(&[]app.User{}).Scan(&users)
	if result.Error != nil {
		return nil,result.Error
	}
	return &users,nil
}

func (u *user) GetById(ctx context.Context, db *gorm.DB, userId string) (*response.User, error) {
	var user response.User
	result := db.WithContext(ctx).First(&app.User{},"id = ?", userId).Scan(&user)
	if result.Error != nil {
		return nil,result.Error
	}
	return &user,nil
}

func (u *user) Post(ctx context.Context, db *gorm.DB, request request.User) (*response.User, error) {
	payload := &app.User{
		Username:     request.Username,
		NamaDepan:    request.NamaDepan,
		NamaBelakang: request.NamaBelakang,
		Level:        request.Level,
		Password:     request.Password,
	}
	result := db.WithContext(ctx).Create(payload)

	if result.Error != nil {
		return nil,result.Error
	}

	return &response.User{
		Id:           payload.Id,
		Username:     payload.Username,
		NamaDepan:    payload.NamaDepan,
		NamaBelakang: payload.NamaBelakang,
		Level:        payload.Level,
	},nil
}

func (u *user) Delete(ctx context.Context, db *gorm.DB, userId string) (bool, error) {
	result := db.WithContext(ctx).Delete(&app.User{}, "id = ?", userId)
	if result.Error != nil {
		return false,result.Error
	}
	return result.RowsAffected > 0, nil
}

func (u *user) Update(ctx context.Context, db *gorm.DB, request request.User) (bool, error) {
	result := db.WithContext(ctx).Where("id = ?",request.Id).Updates(&app.User{
		Username:     request.Username,
		NamaDepan:    request.NamaDepan,
		NamaBelakang: request.NamaBelakang,
		Level:        request.Level,
		Password:     request.Password,
	})
	if result.Error != nil {
		return false,result.Error
	}
	return result.RowsAffected > 0, nil
}
