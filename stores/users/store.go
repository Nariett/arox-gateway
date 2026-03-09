package users

import "github.com/jmoiron/sqlx"

//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE

type Store interface {
	Get(login, password string) (string, error)
}

func (s *store) Get(login, password string) (string, error) {
	return "test", nil
}

type store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return &store{db: db}
}
