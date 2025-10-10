package repopostgres

import (
	"context"

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

func (r *Group) FindByGroupFilter(ctx context.Context, filter *domain.GroupFilter) ([]domain.Group, error) {
	filters := make(map[string]any)
	if len(filter.ReleasePlanIDs) > 0 {
		filters["release_plan_id"] = filter.ReleasePlanIDs
	}
	if len(filter.GroupIDs) > 0 {
		filters["id"] = filter.GroupIDs
	}
	return r.BaseRepository.FindByFilter(ctx, filters)
}
