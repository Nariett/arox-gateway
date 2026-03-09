package rest

import (
	"arox-gateway/rest/endpoints"
	"net/http"

	"github.com/Nariett/arox-pkg/config"
	"github.com/Nariett/arox-pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(cfg *config.Config, endpoints endpoints.Endpoints) chi.Router {
	r := chi.NewRouter()

	r.Route("/token", func(r chi.Router) {
		r.Post("/", endpoints.Token().GetToken())
	})

	r.Route("/products", func(r chi.Router) {
		r.Get("/", endpoints.Products().GetProducts())
		r.Get("/{id}", endpoints.Products().GetProduct())
	})

	r.Route("/categories", func(r chi.Router) {
		r.Get("/", endpoints.Categories().GetCategories())
		r.Get("/{id}", endpoints.Categories().GetCategory())
	})

	r.Group(func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return middleware.AuthMiddleware(next, cfg.Secret)
		})

		r.Route("/test", func(r chi.Router) {
			r.Get("/{id}", endpoints.Categories().GetCategory())
		})
	})
	return r
}
