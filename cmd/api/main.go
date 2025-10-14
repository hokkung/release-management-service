// @title Release Management Service
// @version v1.0.0
// @schemes http https
// @BasePath /
package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

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
	defer func() {
		fmt.Println("gracefully shutdown server")
	}()

	ctx, stopFunc := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopFunc()

	cfg := config.New()
	// db, err := reposqlite.New(*cfg)
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

	srvChanErr := make(chan error, 1)
	go func() {
		if err := server.Start(); err != nil {
			srvChanErr <- err
		}
	}()

	select {
	case srvErr := <-srvChanErr:
		panic(fmt.Errorf("unable to start server: %w", srvErr))
	case <-ctx.Done():
		// do nothing
	}
	if err = server.Stop(); err != nil {
		panic(fmt.Errorf("unable to stop server: %w", err))
	}
}
