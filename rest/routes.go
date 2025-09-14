package rest

import (
	"arox-gateway/rest/endpoints"
	"github.com/go-chi/chi/v5"
)

func NewRouter(endpoints endpoints.Endpoints) chi.Router {
	r := chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Get("/", endpoints.Products().GetProducts())
		r.Get("/{id}", endpoints.Products().GetProduct())
	})

	r.Route("/categories", func(r chi.Router) {
		r.Get("/", endpoints.Categories().GetCategories())
		r.Get("/{id}", endpoints.Categories().GetCategory())
	})

	return r
}
