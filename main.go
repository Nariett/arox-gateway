package main

import (
	"arox-gateway/rest"
	"arox-gateway/rest/endpoints"
	"arox-gateway/rpc"
	"arox-gateway/services"
	"arox-gateway/stores"
	"github.com/Nariett/arox-pkg/config"
	"github.com/Nariett/arox-pkg/db"
	"go.uber.org/fx"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)

		return
	}

	application := fx.New(
		fx.Supply(cfg),
		fx.Provide(
			endpoints.NewHandler,
			rest.NewRouter,
			rpc.ListenServer,
			db.NewPostgres,
		),
		stores.Construct(),
		services.Construct(),
		endpoints.Provider,
		fx.Invoke(
			rest.StartHttpServer,
		),
	)
	application.Run()
}
