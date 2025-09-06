package rpc

import (
	"fmt"
	"github.com/Nariett/arox-pkg/config"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"google.golang.org/grpc"
	"log"
)

func ListenServer(cfg *config.Config) proto.ProductsServiceClient {
	connStr := fmt.Sprintf("%s:%s", cfg.Host, cfg.ProductsPort)
	conn, err := grpc.Dial(connStr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error with grpc connection: %v", err)
	}
	client := proto.NewProductsServiceClient(conn)

	return client
}
