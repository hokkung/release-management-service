package repopostgres

import (
	"context"

	"github.com/hokkung/release-management-service/internal/domain"
	"github.com/hokkung/release-management-service/pkg/gorem"
	"gorm.io/gorm"
)

type Repository struct {
	*gorem.BaseRepository[domain.Repository]
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		BaseRepository: gorem.NewBaseRepository[domain.Repository](db),
	}
}

func (r *Repository) FindActive(ctx context.Context) ([]domain.Repository, error) {
	return gorm.G[domain.Repository](r.GetDB(ctx)).Where("status = ?", "ACTIVE").Find(ctx)
}
