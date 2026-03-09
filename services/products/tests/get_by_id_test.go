package tests

import (
	"arox-gateway/services/products"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	mockproducts "github.com/Nariett/arox-pkg/grpc/pb/mock/products"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
)

var (
	testProduct = &proto.GetProductResponse{
		Product: &proto.Product{
			Id:          1,
			Brand:       "Nike",
			Name:        "T-Shirt HV9803-110",
			CategoryId:  1,
			Price:       145,
			Description: "T-Shirt HV9803-110",
			Sizes: []*proto.Size{
				{
					Size:  "S",
					Count: 11,
				},
				{
					Size:  "M",
					Count: 6,
				},
				{
					Size:  "L",
					Count: 10,
				},
			},
			IsActive:  true,
			CreatedAt: timestamppb.New(time.Now()),
			Images: []*proto.Image{
				{
					Id:        1,
					IdProduct: 1,
					Url:       "http://example.com/image1.png",
					IsMain:    true,
				},
				{
					Id:        2,
					IdProduct: 1,
					Url:       "http://example.com/image2.png",
					IsMain:    false,
				},
			},
		},
	}
)

func TestGetById(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mockproducts.NewMockProductsServiceClient(ctrl)

	t.Run("success", func(t *testing.T) {
		mockClient.EXPECT().GetProduct(ctx, &proto.GetProductRequest{
			Id: int64(1),
		}).Return(testProduct, nil)

		s := products.NewService(mockClient)

		value, err := s.GetById(ctx, 1)
		assert.NoError(t, err)

		assert.Equal(t, testProduct.Product.Id, value.Id)
		assert.Equal(t, testProduct.Product.Brand, value.Brand)
		assert.Equal(t, testProduct.Product.Name, value.Name)
		assert.Equal(t, testProduct.Product.CategoryId, value.CategoryId)
		assert.Equal(t, testProduct.Product.Price, value.Price)
		assert.Equal(t, &testProduct.Product.Description, value.Description)
		assert.Equal(t, len(testProduct.Product.Sizes), len(value.Sizes))
		assert.Equal(t, testProduct.Product.IsActive, value.IsActive)
		assert.Equal(t, testProduct.Product.CreatedAt.AsTime(), value.CreatedAt)
	})

	t.Run("product not found", func(t *testing.T) {
		mockClient.EXPECT().GetProduct(ctx, &proto.GetProductRequest{
			Id: int64(21),
		}).Return(nil, status.Error(codes.NotFound, fmt.Sprintf("product with id %d not found", 21)))

		s := products.NewService(mockClient)

		value, err := s.GetById(ctx, 21)
		assert.Error(t, err)
		assert.Nil(t, value)
	})

	t.Run("error", func(t *testing.T) {
		mockClient.EXPECT().GetProduct(ctx, &proto.GetProductRequest{
			Id: int64(62),
		}).Return(nil, status.Error(codes.Internal, fmt.Sprint("unknow error")))

		s := products.NewService(mockClient)

		value, err := s.GetById(ctx, 62)
		assert.Error(t, err)
		assert.Nil(t, value)
	})
}
