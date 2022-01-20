package setup

import (
	"crud-service/controller"
	"crud-service/repository"
	service2 "crud-service/service"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func setupController(db *gorm.DB, valid *validator.Validate) controller.User {
	repo := repository.NewUser()
	service := service2.NewUser(db,valid,repo)
	controller := controller.NewUser(service)
	return controller
}
