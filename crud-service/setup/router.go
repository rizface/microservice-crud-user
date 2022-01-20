package setup

import (
	myMiddleware "crud-service/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Router(r *chi.Mux,db *gorm.DB, valid *validator.Validate) *chi.Mux {
	controller := setupController(db,valid)

	r.Route("/users", func(r chi.Router) {
		r.Use(myMiddleware.Admin)
		r.Get("/", controller.Get)
		r.Post("/",controller.Post)
		r.Get("/{userId}",controller.GetById)
		r.Delete("/{userId}",controller.Delete)
		r.Put("/{userId}",controller.Update)
	})

	r.Route("/profile", func(r chi.Router) {
		r.Use(myMiddleware.General)
		r.Get("/",controller.Profile)
	})

	return r
}
