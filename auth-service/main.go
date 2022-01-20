package main

import (
	"auth-service/exception"
	"auth-service/helper"
	customMiddleware "auth-service/middleware"
	"auth-service/setup"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	// Load Config(.env)
	//helper.LoadConfig(".env")

	// Create Database
	helper.CreateDatabase()

	// Setup Database
	db := helper.NewConnection()

	// Clean Database Before Input Initial Data & Setup Database Extension
	helper.SetupDatabase(db)

	// Setup Router
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(customMiddleware.Error)

	setup.Router(r,db)


	// Start Server
	err := http.ListenAndServe(":9001", r)
	exception.InternalServerError(err)
}