package helper

import (
	"crud-service/exception"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

func ValidateToken(tokenString string) jwt.MapClaims {
	secret := os.Getenv("JWT_SECRET")
	claims := &jwt.MapClaims{}
	tkn,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret),nil
	})

	fmt.Println(tkn.Valid,err)

	exception.Unauthorized(tkn.Valid == false, "token kamu tidak valid")
	if err != nil {
		exception.Unauthorized(true, err.Error())
	}

	return *claims
}
