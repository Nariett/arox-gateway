package categories

import (
	"arox-gateway/services/categories"
	"net/http"
)

type Endpoint interface {
	GetCategories() http.HandlerFunc
	GetCategory() http.HandlerFunc
}

type endpoint struct {
	srv categories.Service
}

func NewEndpoint(srv categories.Service) Endpoint {
	return &endpoint{
		srv: srv,
	}
}
