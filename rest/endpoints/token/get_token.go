package token

import (
	"encoding/json"
	"github.com/Nariett/arox-pkg/response"
	"net/http"
)

type GetTokenResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (e *endpoint) GetToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.BadRequest(w, err.Error())

			return
		}

		token, err := e.srv.GenerateToken(req.Login, req.Password)
		if err != nil {
			response.BadRequest(w, err.Error())

			return
		}

		response.Ok(w, &GetTokenResponse{Token: token})
	}
}
