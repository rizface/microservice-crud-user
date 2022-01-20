package main

import (
	"crud-service/exception"
	"crud-service/helper"
	myMiddleware "crud-service/middleware"
	"crud-service/setup"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func main() {
	// Load Config
	//helper.LoadConfig(".env")

	// Init Router
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(myMiddleware.Error)

	// Open Connection
	db := helper.NewConnection()

	// Init Validator
	valid := validator.New()

	// Setup Router
	setup.Router(r,db,valid)

	// Start Server
	err := http.ListenAndServe(":9002",r)
	exception.InternalServerError(err)
}