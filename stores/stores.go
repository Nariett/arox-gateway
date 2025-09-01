package stores

import "arox-gateway/stores/users"

type Stores interface {
	Users() users.Store
}

type stores struct {
	users users.Store
}

func New(u users.Store) Stores {
	return &stores{
		users: u,
	}
}

func (s *stores) Users() users.Store { return s.users }
