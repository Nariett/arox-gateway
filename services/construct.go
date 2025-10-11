package services

import (
	"arox-gateway/services/categories"
	"arox-gateway/services/products"
	"arox-gateway/services/token"
	"go.uber.org/fx"
)

func Construct() fx.Option {
	return fx.Provide(
		products.NewService,
		categories.NewService,
		token.NewService,
	)
}
