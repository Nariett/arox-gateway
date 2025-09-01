package stores

import (
	"arox-gateway/stores/users"
	"go.uber.org/fx"
)

func Construct() fx.Option {
	return fx.Provide(
		New,
		users.NewStore,
	)
}
