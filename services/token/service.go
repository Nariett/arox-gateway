package token

import (
	"arox-gateway/stores"
	"github.com/Nariett/arox-pkg/config"
)

type Service interface {
	GenerateToken(login, password string) (token string, err error)
}

type service struct {
	cfg    *config.Config
	stores stores.Stores
}

func NewService(cfg *config.Config, stores stores.Stores) Service {
	return &service{cfg: cfg, stores: stores}
}
