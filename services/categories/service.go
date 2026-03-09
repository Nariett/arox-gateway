package categories

import (
	"arox-gateway/models"
	"context"

	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE

type Service interface {
	GetById(ctx context.Context, id int64) (*models.Category, error)
	List(ctx context.Context) ([]*models.Category, error)
}

type service struct {
	client proto.ProductsServiceClient
}

func NewService(client proto.ProductsServiceClient) Service {
	return &service{client}
}
