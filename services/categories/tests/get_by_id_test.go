package tests

import (
	"arox-gateway/services/categories"
	"context"
	"fmt"
	"testing"

	"github.com/Nariett/arox-pkg/grpc/pb/products"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mockproducts "github.com/Nariett/arox-pkg/grpc/pb/mock/products"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
)

var (
	testCategory = &products.Category{
		Id:   1,
		Name: "Футболки",
		Slug: "t-shirt",
	}
)

func TestGetById(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mockproducts.NewMockProductsServiceClient(ctrl)

	t.Run("success", func(t *testing.T) {
		mockClient.EXPECT().GetCategory(ctx, &proto.GetCategoryRequest{
			Id: int64(1),
		}).Return(&products.GetCategoryResponse{
			Category: testCategory}, nil)

		s := categories.NewService(mockClient)

		value, err := s.GetById(ctx, 1)
		assert.NoError(t, err)

		assert.Equal(t, testCategory.Id, value.Id)
		assert.Equal(t, testCategory.Name, value.Name)
		assert.Equal(t, testCategory.Slug, value.Slug)
	})

	t.Run("category not found", func(t *testing.T) {
		mockClient.EXPECT().GetCategory(ctx, &proto.GetCategoryRequest{
			Id: int64(66),
		}).Return(&products.GetCategoryResponse{
			Category: nil}, status.Error(codes.NotFound, fmt.Sprintf("category with id %d not found", 66)))

		s := categories.NewService(mockClient)

		value, err := s.GetById(ctx, 66)
		assert.Error(t, err)
		assert.Nil(t, value)
	})

	t.Run("error", func(t *testing.T) {
		mockClient.EXPECT().GetCategory(ctx, &proto.GetCategoryRequest{
			Id: int64(66),
		}).Return(&products.GetCategoryResponse{
			Category: nil}, status.Error(codes.Internal, fmt.Sprint("unknow error")))

		s := categories.NewService(mockClient)

		value, err := s.GetById(ctx, 66)
		assert.Error(t, err)
		assert.Nil(t, value)
	})
}
