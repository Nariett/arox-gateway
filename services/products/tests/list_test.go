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
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	mockproducts "github.com/Nariett/arox-pkg/grpc/pb/mock/products"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
)

var (
	testListProducts = &proto.ListProductsResponse{
		Products: []*proto.Product{
			{
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
			{
				Id:          2,
				Brand:       "Adidas",
				Name:        "T-Shirt DB1241-223",
				CategoryId:  1,
				Price:       145,
				Description: "T-Shirt DB1241-223",
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
		},
	}
)

func TestList(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mockproducts.NewMockProductsServiceClient(ctrl)

	t.Run("success", func(t *testing.T) {
		mockClient.EXPECT().ListProducts(ctx, &emptypb.Empty{}).Return(testListProducts, nil)

		s := products.NewService(mockClient)

		value, err := s.List(ctx)
		assert.NoError(t, err)

		for i, product := range value {
			assert.Equal(t, testListProducts.Products[i].Id, product.Id)
			assert.Equal(t, testListProducts.Products[i].Brand, product.Brand)
			assert.Equal(t, testListProducts.Products[i].Name, product.Name)
			assert.Equal(t, testListProducts.Products[i].CategoryId, product.CategoryId)
			assert.Equal(t, testListProducts.Products[i].Price, product.Price)
			assert.Equal(t, &testListProducts.Products[i].Description, product.Description)
			assert.Equal(t, len(testListProducts.Products[i].Sizes), len(product.Sizes))
			assert.Equal(t, testListProducts.Products[i].IsActive, product.IsActive)
			assert.Equal(t, testListProducts.Products[i].CreatedAt.AsTime(), product.CreatedAt)
		}

	})

	t.Run("product not found", func(t *testing.T) {
		mockClient.EXPECT().ListProducts(ctx, &emptypb.Empty{}).Return(nil, status.Error(codes.NotFound, "products not found"))

		s := products.NewService(mockClient)

		value, err := s.List(ctx)
		assert.Error(t, err)
		assert.Nil(t, value)
	})

	t.Run("error", func(t *testing.T) {
		mockClient.EXPECT().ListProducts(ctx, &emptypb.Empty{}).Return(nil, status.Error(codes.Internal, fmt.Sprint("unknow error")))

		s := products.NewService(mockClient)

		value, err := s.List(ctx)
		assert.Error(t, err)
		assert.Nil(t, value)
	})
}
