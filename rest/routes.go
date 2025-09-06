package rest

import (
	"arox-gateway/rest/endpoints"
	"github.com/go-chi/chi/v5"
)

func NewRouter(endpoints endpoints.Endpoints) chi.Router {
	r := chi.NewRouter()

	r.Get("/products/", endpoints.Products().GetProducts())
	r.Get("/product/{id}", endpoints.Products().GetProduct())

	r.Get("/categories/", endpoints.Categories().GetCategories())

	return r
}
