package repopostgres

import (
	"context"
	"errors"

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

func (r *Repository) FindByKey(ctx context.Context, key interface{}) (*domain.Repository, bool, error) {
	ent, err := gorm.G[domain.Repository](r.db).Where("id = ?", key).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &ent, true, nil
}

func (r *Repository) FindByName(ctx context.Context, name string) (*domain.Repository, bool, error) {
	ent, err := gorm.G[domain.Repository](r.db).Where("name = ?", name).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &ent, true, nil
}