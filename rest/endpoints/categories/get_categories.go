package categories

import (
	"net/http"

	"github.com/Nariett/arox-pkg/response"
)

func (e *endpoint) GetCategories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := e.srv.List(r.Context())
		if err != nil {
			response.NotFound(w, "categories not found")

			return
		}

		response.Ok(w, categories)
	}
}
