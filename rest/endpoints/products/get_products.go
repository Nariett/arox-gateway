package products

import (
	"github.com/Nariett/arox-pkg/response"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

func (e *endpoint) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := e.stores.Users().Get("test")
		if err != nil {
			response.NotFound(w, "users not found")

			return
		}

		products, err := e.client.ListProducts(r.Context(), &emptypb.Empty{})
		if err != nil {
			response.NotFound(w, "products not found")

			return
		}

		response.Ok(w, products)
	}
}
