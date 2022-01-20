package myMiddleware

import (
	"crud-service/exception"
	"crud-service/helper"
	"net/http"
)

func Error(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				e := err.(exception.Exception)
				helper.JsonWriter(w,e.Code,e.Error(),nil)
			}
		}()
		w.Header().Set("Content-Type","application/json")
		next.ServeHTTP(w,r)
	})
}
