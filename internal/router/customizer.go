package router

import (
	"github.com/gofiber/swagger"
	"github.com/hokkung/release-management-service/config"
	"github.com/hokkung/release-management-service/internal/delivery/rest/handler"
	server "github.com/hokkung/release-management-service/pkg/srv"
)

// Customizer ...
type Customizer struct {
	cfg               config.Configuration
	repositoryHandler *handler.Repository
}

// NewCustomizer creates instance
func NewCustomizer(
	cfg config.Configuration,
	repositoryHandler *handler.Repository,
) *Customizer {
	return &Customizer{
		cfg:               cfg,
		repositoryHandler: repositoryHandler,
	}
}

func (c *Customizer) Register(srv *server.Server) {
	srv.App.Get("/swagger/*", swagger.HandlerDefault)

	v1 := srv.App.Group("/api/v1")
	repositoryGroup := v1.Group("/repositories")
	repositoryGroup.Get("", c.repositoryHandler.List)
	repositoryGroup.Post("/sync", c.repositoryHandler.Sync)
	repositoryGroup.Post("/register", c.repositoryHandler.Register)
}
