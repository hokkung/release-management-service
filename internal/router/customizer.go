package router

import (
	"github.com/gofiber/swagger"
	"github.com/hokkung/release-management-service/config"
	server "github.com/hokkung/release-management-service/pkg/srv"
)

// Customizer ...
type Customizer struct {
	cfg config.Configuration
}

// NewCustomizer creates instance
func NewCustomizer(
	cfg config.Configuration,
) *Customizer {
	return &Customizer{
		cfg: cfg,
	}
}

func (c *Customizer) Register(srv *server.Server) {
	srv.App.Get("/swagger/*", swagger.HandlerDefault)
}
