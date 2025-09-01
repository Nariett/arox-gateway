package endpoints

import (
	"arox-gateway/rest/endpoints/product"
	"arox-gateway/stores"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"go.uber.org/fx"
)

var Provider = fx.Provide(
	product.NewEndpoint,
)

type Endpoints interface {
	Product() product.Endpoint
}

type endpoints struct {
	client  proto.ProductsServiceClient
	product product.Endpoint
	stores  stores.Stores
}

func NewHandler(client proto.ProductsServiceClient, product product.Endpoint, stores stores.Stores) Endpoints {
	return &endpoints{
		product: product,
		client:  client,
		stores:  stores,
	}
}

func (e *endpoints) Product() product.Endpoint { return e.product }
