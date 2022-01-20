package setup

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Router(r *chi.Mux, db *gorm.DB) *chi.Mux {
	controller := setController(db)
	r.Post("/login",controller.Login)
	return r
}
