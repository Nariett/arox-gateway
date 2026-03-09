package tests

import (
	"arox-gateway/services/categories"
	"context"
	"errors"
	"testing"

	"github.com/Nariett/arox-pkg/grpc/pb/products"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mockproducts "github.com/Nariett/arox-pkg/grpc/pb/mock/products"
)

var (
	testCategories = []*products.Category{
		{
			Id:   1,
			Name: "Футболки",
			Slug: "t-shirt",
		},
		{
			Id:   2,
			Name: "Зип-худи",
			Slug: "zip-hoodie",
		},
		{
			Id:   3,
			Name: "Лонгсливы",
			Slug: "longsleeve",
		},
		{
			Id:   4,
			Name: "Шорты",
			Slug: "shorts",
		},
		{
			Id:   5,
			Name: "Джинсы",
			Slug: "jeans",
		},
		{
			Id:   6,
			Name: "Штаны",
			Slug: "pants",
		},
		{
			Id:   7,
			Name: "Брюки",
			Slug: "trousers",
		},
		{
			Id:   8,
			Name: "Худи",
			Slug: "hoodie",
		},
		{
			Id:   9,
			Name: "Куртка",
			Slug: "jacket",
		},
		{
			Id:   10,
			Name: "Нижнее белье",
			Slug: "underwear",
		},
		{
			Id:   11,
			Name: "Кроссовки",
			Slug: "sneakers",
		},
		{
			Id:   12,
			Name: "Ботинки",
			Slug: "boots",
		},
		{
			Id:   13,
			Name: "Балаклавы",
			Slug: "balaclava",
		},
	}
)

func TestList(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mockproducts.NewMockProductsServiceClient(ctrl)

	t.Run("success", func(t *testing.T) {
		mockClient.EXPECT().ListCategories(ctx, gomock.Any()).Return(&products.ListCategoriesResponse{
			Categories: testCategories}, nil)

		s := categories.NewService(mockClient)

		value, err := s.List(ctx)
		assert.NoError(t, err)

		assert.Len(t, value, len(testCategories))
		for i, v := range value {
			assert.Equal(t, testCategories[i].Id, v.Id)
			assert.Equal(t, testCategories[i].Name, v.Name)
			assert.Equal(t, testCategories[i].Slug, v.Slug)
		}
	})

	t.Run("nil response", func(t *testing.T) {
		mockClient.EXPECT().ListCategories(ctx, gomock.Any()).Return(&products.ListCategoriesResponse{
			Categories: nil}, nil)

		s := categories.NewService(mockClient)

		value, err := s.List(ctx)
		assert.NoError(t, err)
		assert.Len(t, value, 0)
	})

	t.Run("error", func(t *testing.T) {
		mockClient.EXPECT().ListCategories(ctx, gomock.Any()).Return(nil, errors.New("error"))

		s := categories.NewService(mockClient)

		value, err := s.List(ctx)
		assert.Error(t, err)
		assert.Nil(t, value)
	})
}
