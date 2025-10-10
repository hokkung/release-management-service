// @title Release Management Service
// @version v1.0.0
// @schemes http https
// @BasePath /
package main

import (
	"context"

	"github.com/google/go-github/v75/github"
	"github.com/hokkung/release-management-service/config"
	_ "github.com/hokkung/release-management-service/docs"
	repopostgres "github.com/hokkung/release-management-service/internal/repository/postgres"
	"github.com/hokkung/release-management-service/internal/router"
	"github.com/hokkung/release-management-service/internal/service/group_item"
	"github.com/hokkung/release-management-service/internal/service/release_plan"
	"github.com/hokkung/release-management-service/internal/service/repository"
	"github.com/hokkung/release-management-service/pkg/githuby"
	"github.com/hokkung/release-management-service/pkg/srv"
)

func main() {
	cfg := config.New()
	ctx := context.Background()
	db, err := repopostgres.New(*cfg)
	if err != nil {
		panic(err)
	}

	githubClient := github.NewClient(nil).WithAuthToken("")
	githubService := githuby.New(githubClient)
	reporepo := repopostgres.NewRepository(db)

	groupRepository := repopostgres.NewGroupItem(db)
	groupItemService := group_item.NewGroupItem(groupRepository)
	releasePlanRepository := repopostgres.NewReleasePlan(db)
	releasePlanService := release_plan.NewReleasePlan(releasePlanRepository)
	repoService := repository.NewRepository(reporepo, githubService, groupItemService, releasePlanService)
	// err = repoService.Register(ctx, &service.RegisterRequest{
	// 	Name:  "go-groceries",
	// 	Owner: "hokkung",
	// })
	// if err != nil {
	// 	panic(err)
	// }
	err = repoService.Sync(ctx, &repository.SyncRequest{
		RepositoryName: "go-groceries",
	})
	if err != nil {
		panic(err)
	}

	customizer := router.NewCustomizer(*cfg)
	server := srv.New(customizer)
	err = server.Start()
	if err != nil {
		panic(err)
	}
}
