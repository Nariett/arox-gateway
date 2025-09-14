package categories

import (
	"arox-gateway/models"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *service) List(ctx context.Context) ([]*models.Category, error) {
	resp, err := s.client.ListCategories(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("error with rpc ListCategories: %w", err)
	}

	var categories []*models.Category

	for _, category := range resp.Categories {
		categories = append(categories, &models.Category{
			Id:   category.Id,
			Name: category.Name,
			Slug: category.Slug,
		})
	}

	return categories, nil
}
