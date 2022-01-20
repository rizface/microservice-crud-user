package customMiddleware

import (
	"auth-service/exception"
	"auth-service/helper"
	"net/http"
)

func Error(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			e := err.(exception.Exception)
			helper.JsonWriter(w,e.Code,e.Error(),nil)
		}()
		w.Header().Add("Content-Type","application/json")
		next.ServeHTTP(w,r)
	})
}
