package products

import (
	"arox-gateway/models"
	"context"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
)

type Service interface {
	GetById(ctx context.Context, id int64) (*models.Product, error)
	List(ctx context.Context) ([]*models.Product, error)
}

type service struct {
	client proto.ProductsServiceClient
}

func NewService(client proto.ProductsServiceClient) Service {
	return &service{client: client}
}
