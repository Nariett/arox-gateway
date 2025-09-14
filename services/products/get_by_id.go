package products

import (
	"arox-gateway/models"
	"context"
	"fmt"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
)

func (s *service) GetById(ctx context.Context, id int64) (*models.Product, error) {
	resp, err := s.client.GetProduct(ctx, &proto.GetProductRequest{
		Id: id,
	})

	if err != nil {
		return nil, fmt.Errorf("error with rpc GetProduct: %w", err)
	}

	var sizes []models.Size

	for _, sz := range resp.Product.Sizes {
		sizes = append(sizes, models.Size{
			Size:  sz.Size,
			Count: sz.Count,
		})
	}

	var images []models.Image

	for _, img := range resp.Product.Images {
		images = append(images, models.Image{
			Id:        img.Id,
			IdProduct: img.IdProduct,
			Url:       img.Url,
			IsMain:    img.IsMain,
			IsActive:  img.IsActive,
		})
	}

	product := &models.Product{
		Id:          resp.Product.Id,
		Brand:       resp.Product.Brand,
		Name:        resp.Product.Name,
		CategoryId:  resp.Product.CategoryId,
		Price:       resp.Product.Price,
		Description: &resp.Product.Description,
		IsActive:    resp.Product.IsActive,
		CreatedAt:   resp.Product.CreatedAt.AsTime(),
		Sizes:       sizes,
		Images:      images,
	}

	return product, nil
}
