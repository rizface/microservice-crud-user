package controller

import "net/http"

type Auth interface {
	Login(w http.ResponseWriter, r *http.Request)
}
