package repopostgres

import (
	"github.com/hokkung/release-management-service/internal/domain"
	"github.com/hokkung/release-management-service/pkg/gorem"
	"gorm.io/gorm"
)

type Group struct {
	*gorem.BaseRepository[domain.Group]
}

func NewGroup(db *gorm.DB) *Group {
	return &Group{
		BaseRepository: gorem.NewBaseRepository[domain.Group](db),
	}
}
