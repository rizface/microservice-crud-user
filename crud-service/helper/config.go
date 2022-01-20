package helper

import (
	"crud-service/exception"
	"github.com/joho/godotenv"
)

func LoadConfig(filenames ...string) {
	err := godotenv.Load(filenames...)
	exception.InternalServerError(err)
}
