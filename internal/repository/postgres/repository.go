package repopostgres

import (
	"context"

	"github.com/hokkung/release-management-service/internal/domain"
	"gorm.io/gorm"
)

type GormKey string

const GormContext GormKey = "gormContext"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, ent *domain.Repository) error {
	// db := r.getDB(ctx)
	err := gorm.G[domain.Repository](r.db).Create(ctx, ent)
	if err != nil {
		return nil
	}
	return nil
}

func (r *Repository) getDB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(GormContext).(*gorm.DB)
	if !ok {
		return r.db
	}
	return tx
}
