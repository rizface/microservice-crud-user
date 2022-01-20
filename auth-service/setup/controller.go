package setup

import (
	"auth-service/controller"
	"auth-service/repository"
	service2 "auth-service/service"
	"gorm.io/gorm"
)

func setController(db *gorm.DB) controller.Auth {
	repo := repository.NewUser()
	service := service2.NewAuth(db,repo)
	controller := controller.NewAuth(service)
	return controller
}
