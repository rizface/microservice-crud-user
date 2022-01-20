package exception

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func InternalServerError(err error) {
	if err != nil {
		panic(Exception{
			Code: http.StatusInternalServerError,
			err:  "terjadi kesalahan pada sistem kami",
		})
	}
}

func NotFound(err error, msg string) {
	if err != nil {
		panic(Exception{
			Code: http.StatusNotFound,
			err:  msg,
		})
	}
}

func BadRequest(err error, msg string) {
	if err != nil {
		panic(Exception{
			Code: http.StatusBadRequest,
			err:  msg,
		})
	}
}

func Unauthorized(cond bool, msg string) {
	if cond == true {
		panic(Exception{
			Code: http.StatusUnauthorized,
			err:  msg,
		})
	}
}


func Validate(err error) {
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("%s%s",v.Field(),msg[v.Tag()])
			BadRequest(errors.New(msg),msg)
		}
	}
}