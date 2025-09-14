package services

import (
	"arox-gateway/services/categories"
	"arox-gateway/services/products"
	"go.uber.org/fx"
)

func Construct() fx.Option {
	return fx.Provide(
		products.NewService,
		categories.NewService,
	)
}
