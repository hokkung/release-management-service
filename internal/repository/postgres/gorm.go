package repopostgres

import (
	"github.com/hokkung/release-management-service/config"
	"github.com/hokkung/release-management-service/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg config.Configuration) (*gorm.DB, error) {
	dsn := cfg.DB.DSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// TODO: move to migrator
	err = db.AutoMigrate(&domain.Repository{})
	if err != nil {
		return nil, err
	}

	return db, err
}
