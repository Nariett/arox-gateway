package token

import (
	"github.com/Nariett/arox-pkg/config"
	"github.com/google/uuid"
)

type Service interface {
	GenerateToken(uuid uuid.UUID) (token string, err error)
}

type service struct {
	cfg *config.Config
}

func NewService(cfg *config.Config) Service {
	return &service{cfg: cfg}
}
