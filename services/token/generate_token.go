package token

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

func (s *service) GenerateToken(login, password string) (string, error) {
	_, err := s.stores.Users().Get(login, password)
	if err != nil {
		return "", err
	}

	userid := uuid.New()

	claims := jwt.MapClaims{
		"uuid": userid.String(),
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(30 * time.Minute).Unix(),
	}

	secret := s.cfg.Secret
	if secret == "" {
		return "", errors.New("no jwt secret")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
