package endpoint

import (
	"arox-gateway/rest/endpoint/product"
	"go.uber.org/fx"
)

var Provider = fx.Provide(
	product.NewEndpoint,
)

type Endpoints interface {
	Product() product.Endpoint
}

type endpoints struct {
	product product.Endpoint
}

func NewHandler(product product.Endpoint) Endpoints {
	return &endpoints{
		product: product,
	}
}

func (e *endpoints) Product() product.Endpoint { return e.product }
