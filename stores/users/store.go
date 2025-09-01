package users

import "github.com/jmoiron/sqlx"

type Store interface {
	Get(username string) (string, error)
}

func (s *store) Get(username string) (string, error) {
	return username, nil
}

type store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return &store{db: db}
}
