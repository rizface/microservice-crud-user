package helper

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(plain string) string  {
	passByte,_ := bcrypt.GenerateFromPassword([]byte(plain),bcrypt.DefaultCost)
	return string(passByte)
}
