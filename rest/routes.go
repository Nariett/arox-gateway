package rest

import (
	"arox-gateway/rest/endpoint"
	"github.com/go-chi/chi/v5"
)

func NewRouter(endpoints endpoint.Endpoints) chi.Router {
	r := chi.NewRouter()

	r.Get("/products/", endpoints.Product().Gets())

	return r
}
