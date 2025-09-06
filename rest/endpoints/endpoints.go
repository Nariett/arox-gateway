package endpoints

import (
	"arox-gateway/rest/endpoints/categories"
	"arox-gateway/rest/endpoints/products"
	"arox-gateway/stores"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"go.uber.org/fx"
)

var Provider = fx.Provide(
	products.NewEndpoint,
	categories.NewEndpoint,
)

type Endpoints interface {
	Products() products.Endpoint
	Categories() categories.Endpoint
}

type endpoints struct {
	client     proto.ProductsServiceClient
	products   products.Endpoint
	categories categories.Endpoint
	stores     stores.Stores
}

func NewHandler(client proto.ProductsServiceClient, products products.Endpoint, categories categories.Endpoint, stores stores.Stores) Endpoints {
	return &endpoints{
		client:     client,
		products:   products,
		categories: categories,
		stores:     stores,
	}
}

func (e *endpoints) Products() products.Endpoint     { return e.products }
func (e *endpoints) Categories() categories.Endpoint { return e.categories }
