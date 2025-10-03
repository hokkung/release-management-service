package main

import (
	"github.com/hokkung/release-management-service/config"
	repopostgres "github.com/hokkung/release-management-service/internal/repository/postgres"
	"github.com/hokkung/release-management-service/internal/router"
	"github.com/hokkung/release-management-service/internal/service"
	"github.com/hokkung/release-management-service/pkg/srv"
)

func main() {
	cfg := config.New()
	db, err := repopostgres.New(*cfg)
	if err != nil {
		panic(err)
	}

	reporepo := repopostgres.NewRepository(db)
	repoService := service.NewRepository(reporepo)
	repoService.Create()
	
	customizer := router.NewCustomizer(*cfg)
	server := srv.New(customizer)
	err = server.Start()
	if err != nil {
		panic(err)
	}
}
