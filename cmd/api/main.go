// @title Release Management Service
// @version v1.0.0
// @schemes http https
// @BasePath /
package main

import (
	"github.com/google/go-github/v75/github"
	"github.com/hokkung/release-management-service/config"
	_ "github.com/hokkung/release-management-service/docs"
	"github.com/hokkung/release-management-service/internal/delivery/rest/handler"
	repopostgres "github.com/hokkung/release-management-service/internal/repository/postgres"
	"github.com/hokkung/release-management-service/internal/router"
	"github.com/hokkung/release-management-service/internal/service/group"
	"github.com/hokkung/release-management-service/internal/service/release_plan"
	"github.com/hokkung/release-management-service/internal/service/repository"
	"github.com/hokkung/release-management-service/pkg/githuby"
	"github.com/hokkung/release-management-service/pkg/srv"
)

func main() {
	cfg := config.New()
	db, err := repopostgres.New(*cfg)
	if err != nil {
		panic(err)
	}

	githubClient := github.NewClient(nil).WithAuthToken(cfg.GitHub.Token)
	githubService := githuby.New(githubClient)
	reporepo := repopostgres.NewRepository(db)

	groupItemRepository := repopostgres.NewGroupItem(db)
	groupItemService := group.NewGroupItem(groupItemRepository)
	groupRepository := repopostgres.NewGroup(db)
	groupService := group.NewGroup(groupRepository, groupItemService)
	releasePlanRepository := repopostgres.NewReleasePlan(db)
	releasePlanService := release_plan.NewReleasePlan(releasePlanRepository, groupService, groupItemService)
	repoService := repository.NewRepository(reporepo, githubService, groupItemService, releasePlanService)

	releasePlanHandler := handler.NewReleasePlan(releasePlanService)
	repositoryHandler := handler.NewRepository(repoService, *cfg)
	groupItemHandler := handler.NewGroupItem(groupItemService, groupService)
	groupHander := handler.NewGroup(groupService, releasePlanService)
	customizer := router.NewCustomizer(*cfg, repositoryHandler, releasePlanHandler, groupItemHandler, groupHander)
	server := srv.New(customizer)
	err = server.Start()
	if err != nil {
		panic(err)
	}
}
