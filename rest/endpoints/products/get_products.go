package products

import (
	"net/http"

	"github.com/Nariett/arox-pkg/response"
)

func (e *endpoint) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := e.stores.Users().Get("test", "test")
		if err != nil {
			response.NotFound(w, err.Error())

			return
		}

		products, err := e.srv.List(r.Context())
		if err != nil {
			response.NotFound(w, err.Error())

			return
		}

		response.Ok(w, products)
	}
}
