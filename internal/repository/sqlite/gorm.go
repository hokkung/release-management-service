package reposqlite

import (
	"github.com/hokkung/release-management-service/config"
	"github.com/hokkung/release-management-service/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(cfg config.Configuration) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("rms.db"), &gorm.Config{
	})
	if err != nil {
		return nil, err
	}

	// TODO: move to migrator
	ents := []interface{}{
		&domain.Repository{},
		&domain.ReleasePlan{},
		&domain.GroupItem{},
		&domain.Group{},
	}
	for _, ent := range ents {
		err = db.AutoMigrate(ent)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
