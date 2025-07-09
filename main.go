package main

import (
	"arox-gateway/schema"
	"arox-gateway/storage"
	"github.com/Nariett/arox-pkg/config"
	"go.uber.org/fx"
)

func main() {
	application := fx.New(
		fx.Provide(
			config.New,
			storage.NewDBConnection,
		), fx.Invoke(
			schema.Migrate),
	)
	application.Run()
}
