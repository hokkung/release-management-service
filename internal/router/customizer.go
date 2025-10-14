package router

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/swagger"
	"github.com/hokkung/release-management-service/config"
	"github.com/hokkung/release-management-service/internal/delivery/rest/handler"
	server "github.com/hokkung/release-management-service/pkg/srv"
)

// Customizer ...
type Customizer struct {
	cfg                config.Configuration
	repositoryHandler  *handler.Repository
	releasePlanHandler *handler.ReleasePlan
	groupItem          *handler.GroupItem
	groupHandler       *handler.Group
}

// NewCustomizer creates instance
func NewCustomizer(
	cfg config.Configuration,
	repositoryHandler *handler.Repository,
	releasePlanHandler *handler.ReleasePlan,
	groupItem *handler.GroupItem,
	groupHandler *handler.Group,
) *Customizer {
	return &Customizer{
		cfg:                cfg,
		repositoryHandler:  repositoryHandler,
		releasePlanHandler: releasePlanHandler,
		groupItem:          groupItem,
		groupHandler:       groupHandler,
	}
}

func (c *Customizer) Register(srv *server.Server) {
	srv.App.Use(fiberzap.New(fiberzap.Config{
		SkipURIs: []string{"/health"},
	}))

	srv.App.Get("/swagger/*", swagger.HandlerDefault)

	v1 := srv.App.Group("/api/v1")
	repositoryGroup := v1.Group("/repositories")
	repositoryGroup.Get("", c.repositoryHandler.List)
	repositoryGroup.Post("/sync", c.repositoryHandler.Sync)
	repositoryGroup.Post("/register", c.repositoryHandler.Register)

	releasePlanGroup := v1.Group("/release-plans")
	releasePlanGroup.Post("", c.releasePlanHandler.List)
	releasePlanGroup.Post("/:id/update", c.releasePlanHandler.Update)

	groupItemGroup := v1.Group("group-items")
	groupItemGroup.Post("/:id/move", c.groupItem.Move)

	groupGroup := v1.Group("groups")
	groupGroup.Post("", c.groupHandler.CreateGroup)
	groupGroup.Delete("/:id", c.groupHandler.Remove)
	groupGroup.Post("/:id/update-status", c.groupHandler.UpdateStatus)
}
