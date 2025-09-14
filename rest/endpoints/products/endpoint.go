package products

import (
	"arox-gateway/services/products"
	"arox-gateway/stores"
	"net/http"
)

type Endpoint interface {
	GetProducts() http.HandlerFunc
	GetProduct() http.HandlerFunc
}

type endpoint struct {
	stores stores.Stores
	srv    products.Service
}

func NewEndpoint(stores stores.Stores, srv products.Service) Endpoint {
	return &endpoint{
		stores: stores,
		srv:    srv,
	}
}
