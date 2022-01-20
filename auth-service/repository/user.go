package repository

import (
	"auth-service/model/response"
	"context"
	"gorm.io/gorm"
)

type User interface {
	FindByUsername(ctx context.Context, db *gorm.DB,username string) (*response.User,error)
}
