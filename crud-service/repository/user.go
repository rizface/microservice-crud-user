package repository

import (
	"context"
	"crud-service/model/request"
	"crud-service/model/response"
	"gorm.io/gorm"
)

type User interface {
	Get(ctx context.Context,db *gorm.DB) (*[]response.User,error)
	GetById(ctx context.Context, db *gorm.DB,userId string) (*response.User,error)
	Post(ctx context.Context,db *gorm.DB, request request.User) (*response.User,error)
	Delete(ctx context.Context, db *gorm.DB,userId string) (bool,error)
	Update(ctx context.Context, db *gorm.DB, request request.User) (bool,error)
}
