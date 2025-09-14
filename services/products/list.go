package products

import (
	"arox-gateway/models"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *service) List(ctx context.Context) ([]*models.Product, error) {
	resp, err := s.client.ListProducts(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("error with rpc ListProducts: %w", err)
	}

	var products []*models.Product

	for _, p := range resp.Products {
		var sizes []models.Size

		for _, sz := range p.Sizes {
			sizes = append(sizes, models.Size{
				Size:  sz.Size,
				Count: sz.Count,
			})
		}

		var images []models.Image

		for _, img := range p.Images {
			images = append(images, models.Image{
				Id:        img.Id,
				IdProduct: img.IdProduct,
				Url:       img.Url,
				IsMain:    img.IsMain,
				IsActive:  img.IsActive,
			})
		}

		products = append(products, &models.Product{
			Id:          p.Id,
			Brand:       p.Brand,
			Name:        p.Name,
			CategoryId:  p.CategoryId,
			Price:       p.Price,
			Description: &p.Description,
			IsActive:    p.IsActive,
			CreatedAt:   p.CreatedAt.AsTime(),
			Sizes:       sizes,
			Images:      images,
		})

	}

	return products, nil
}
