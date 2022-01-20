package exception

import "net/http"

func InternalServerError(err error) {
	if err != nil {
		panic(Exception{
			Code: http.StatusInternalServerError,
			err:  "terjadi kesalahan pada sistem kami",
		})
	}
}

func Unauthorized(err error, msg string) {
	if err != nil {
		panic(Exception{
			Code: http.StatusUnauthorized,
			err:  msg,
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