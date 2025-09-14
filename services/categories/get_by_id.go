package categories

import (
	"arox-gateway/models"
	"context"
	"fmt"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
)

func (s *service) GetById(ctx context.Context, id int64) (*models.Category, error) {
	category, err := s.client.GetCategory(ctx, &proto.GetCategoryRequest{
		Id: id,
	})
	if err != nil {
		return nil, fmt.Errorf("error with rpc GetCategory: %w", err)
	}

	return &models.Category{
		Id:   category.Category.Id,
		Name: category.Category.Name,
		Slug: category.Category.Slug,
	}, nil
}
