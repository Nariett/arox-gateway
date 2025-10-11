package token

import (
	"github.com/Nariett/arox-pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

type GetTokenResponse struct {
	Token string `json:"token"`
}

func (e *endpoint) GetToken() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		uuidStr := chi.URLParam(req, "uuid")
		if uuidStr == "" {
			response.BadRequest(w, "'uuid' param is required")

			return
		}

		parsedUUID, err := uuid.Parse(uuidStr)
		if err != nil {
			response.BadRequest(w, "invalid uuid")

			return
		}

		token, err := e.srv.GenerateToken(parsedUUID)
		if err != nil {
			response.BadRequest(w, err.Error())

			return
		}

		response.Ok(w, &GetTokenResponse{Token: token})
	}
}
