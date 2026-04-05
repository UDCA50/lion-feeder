package api

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux, db *sql.DB) {
	r.Route("/api", func(r chi.Router) {
		r.Get("/licenses",     ListLicenses(db))
		r.Post("/licenses",    CreateLicense(db))
		r.Put("/licenses/{id}", UpdateLicense(db))
		r.Delete("/licenses/{id}", DeleteLicense(db))
	})
}