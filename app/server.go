package app

import (
	"net/http"

	"github.com/mnabil1718/go-restful-api/middleware"
)

type AddressURL string // needed for dependency injection binding

func NewServer(authMiddleware *middleware.AuthMiddleware, addressUrl AddressURL) *http.Server {
	return &http.Server{
		Addr:    string(addressUrl),
		Handler: authMiddleware,
	}
}
