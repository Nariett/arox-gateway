package product

import (
	"arox-gateway/stores"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"net/http"
)

type Endpoint interface {
	GetProducts() http.HandlerFunc
}

type endpoint struct {
	client proto.ProductsServiceClient
	stores stores.Stores
}

func NewEndpoint(client proto.ProductsServiceClient, stores stores.Stores) Endpoint {
	return &endpoint{
		client: client,
		stores: stores,
	}
}
