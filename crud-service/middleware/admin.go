package myMiddleware

import (
	"crud-service/exception"
	"crud-service/helper"
	"net/http"
	"strings"
)

func Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		exception.Unauthorized(strings.Contains(token, "Bearer") == false, "token kamu tidak valid")
		exception.Unauthorized(len(strings.Split(token, " ")) != 2, "token kamu tidak valid")
		token = strings.Split(token, " ")[1]
		claim := helper.ValidateToken(token)
		if claim["level"] == "admin" {
			next.ServeHTTP(w,r)
		} else {
			exception.Unauthorized(true, "tidak punya akses ke fitur ini")
		}
	})
}
