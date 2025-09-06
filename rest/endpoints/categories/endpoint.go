package categories

import (
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"net/http"
)

type Endpoint interface {
	GetCategories() http.HandlerFunc
}

type endpoint struct {
	client proto.ProductsServiceClient
}

func NewEndpoint(client proto.ProductsServiceClient) Endpoint {
	return &endpoint{
		client: client,
	}
}
