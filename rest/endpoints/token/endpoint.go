package token

import (
	"arox-gateway/services/token"
	"net/http"
)

type Endpoint interface {
	GetToken() http.HandlerFunc
}

type endpoint struct {
	srv token.Service
}

func NewEndpoint(srv token.Service) Endpoint {
	return &endpoint{
		srv: srv,
	}
}
