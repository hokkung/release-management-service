package repopostgres

import (
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
