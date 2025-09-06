package categories

import (
	"github.com/Nariett/arox-pkg/response"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

func (e *endpoint) GetCategories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := e.client.ListCategories(r.Context(), &emptypb.Empty{})
		if err != nil {
			response.NotFound(w, "categories not found")

			return
		}

		response.Ok(w, categories)
	}
}
