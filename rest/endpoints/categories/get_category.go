package categories

import (
	"github.com/Nariett/arox-pkg/response"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

func (e *endpoint) GetCategory() http.HandlerFunc {
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

		category, err := e.srv.GetById(r.Context(), int64(id))
		if err != nil {
			st, ok := status.FromError(err)
			if ok {
				switch st.Code() {
				case codes.NotFound:
					response.NotFound(w, err.Error())
				default:
					response.BadRequest(w, err.Error())
				}

				return
			}

			response.BadRequest(w, err.Error())

			return
		}

		response.Ok(w, category)
	}
}
