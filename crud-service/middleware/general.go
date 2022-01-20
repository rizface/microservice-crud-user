package myMiddleware

import (
	"context"
	"crud-service/exception"
	"crud-service/helper"
	"net/http"
	"strings"
)

func General(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		exception.Unauthorized(strings.Contains(token, "Bearer") == false, "token kamu tidak valid")
		exception.Unauthorized(len(strings.Split(token, " ")) != 2, "token kamu tidak valid")
		token = strings.Split(token, " ")[1]
		claim := helper.ValidateToken(token)
		if claim["level"] == "admin" || claim["level"] == "user"{
			r := r.WithContext(context.WithValue(r.Context(),"userId",claim["id"]))
			next.ServeHTTP(w,r)
		} else {
			exception.Unauthorized(true, "tidak punya akses ke fitur ini")
		}
	})
}
