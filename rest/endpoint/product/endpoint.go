package product

import "net/http"

type Endpoint interface {
	Gets() http.HandlerFunc
}

type endpoint struct {
}

func NewEndpoint() Endpoint {
	return &endpoint{}
}
