package helper

import (
	"auth-service/exception"
	"auth-service/model/response"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type claims struct {
	Id           string `json:"id"`
	Username     string `json:"username"`
	NamaDepan    string `json:"nama_depan"`
	NamaBelakang string `json:"nama_belakang"`
	Level        string `json:"level"`
	jwt.RegisteredClaims
}

func GenerateToken(user response.User) string {
	secret := os.Getenv("JWT_SECRET")
	claim := claims{
		Id:           user.Id,
		Username:     user.Username,
		NamaDepan:    user.NamaDepan,
		NamaBelakang: user.NamaBelakang,
		Level:        user.Level,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "test sejuta cita",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString,err := token.SignedString([]byte(secret))
	exception.Unauthorized(err, "terjadi kesalahan saat membuat kredensial")
	return tokenString
}
