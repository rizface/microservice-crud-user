package helper

import (
	"auth-service/exception"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(plain string) string {
	passByte,err := bcrypt.GenerateFromPassword([]byte(plain),bcrypt.DefaultCost)
	exception.InternalServerError(err)
	return string(passByte)
}

func ComparePassword(plain,hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(plain))
	return err
}
