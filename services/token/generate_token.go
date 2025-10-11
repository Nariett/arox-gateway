package token

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

func (s *service) GenerateToken(uuid uuid.UUID) (string, error) {
	t := time.Now().Add(30 * time.Minute)
	exp := t.Unix()
	claims := jwt.MapClaims{
		"uuid": uuid.String(),
		"exp":  exp,
	}

	secret := s.cfg.Secret
	if secret == "" {
		return "", errors.New("no jwt secret")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
