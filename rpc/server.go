package rpc

import (
	"context"
	"fmt"
	"log"

	"github.com/Nariett/arox-pkg/config"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ListenServer(lc fx.Lifecycle, cfg *config.Config) (proto.ProductsServiceClient, error) {
	connStr := fmt.Sprintf("%s:%s", cfg.Host, cfg.ProductsPort)

	conn, err := grpc.NewClient(
		connStr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, fmt.Errorf("grpc dial %s: %w", connStr, err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Println("Closing gRPC connection")

			return conn.Close()
		},
	})

	return proto.NewProductsServiceClient(conn), nil
}
