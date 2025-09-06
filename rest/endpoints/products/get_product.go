package products

import (
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"github.com/Nariett/arox-pkg/response"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (e *endpoint) GetProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idQuery := chi.URLParam(r, "id")
		if idQuery == "" {
			response.BadRequest(w, "id is required")

			return
		}

		id, err := strconv.ParseUint(idQuery, 10, 64)
		if err != nil {
			response.BadRequest(w, err.Error())

			return
		}

		products, err := e.client.GetProduct(r.Context(), &proto.GetProductRequest{
			Id: int64(id),
		})
		if err != nil {
			response.NotFound(w, "products not found")

			return
		}

		response.Ok(w, products)
	}
}
