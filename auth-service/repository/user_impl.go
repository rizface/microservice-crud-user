package repository

import (
	"auth-service/model/app"
	"auth-service/model/response"
	"context"
	"gorm.io/gorm"
)

type user struct{}

func NewUser() User{
	return &user{}
}

func (u user) FindByUsername(ctx context.Context, db *gorm.DB,username string) (*response.User,error) {
	var user response.User
	result := db.WithContext(ctx).First(&app.User{},"username = ?", username).Scan(&user)
	if result.Error != nil {
		return nil,result.Error
	}
	return &user,nil
}
