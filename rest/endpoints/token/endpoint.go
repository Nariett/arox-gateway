package token

import (
	"arox-gateway/services/token"
	"net/http"

	"github.com/Nariett/arox-pkg/config"

	"github.com/Nariett/arox-pkg/api/integrations/minio"
)

type Endpoint interface {
	GetToken() http.HandlerFunc
}

type endpoint struct {
	cfg   *config.Config
	srv   token.Service
	minio minio.Minio
}

func NewEndpoint(cfg *config.Config, srv token.Service, minio minio.Minio) Endpoint {
	return &endpoint{
		cfg:   cfg,
		srv:   srv,
		minio: minio,
	}
}
