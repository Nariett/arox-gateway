package main

import (
	"arox-gateway/rest"
	"arox-gateway/rest/endpoint"
	"arox-gateway/storage"
	"github.com/Nariett/arox-pkg/config"
	"go.uber.org/fx"
)

func main() {
	application := fx.New(
		fx.Provide(
			config.New,
			storage.NewDBConnection,
			endpoint.NewHandler,
			rest.NewRouter,
		),
		endpoint.Provider,
		fx.Invoke(
			rest.StartHttpServer,
		),
	)
	application.Run()
}
