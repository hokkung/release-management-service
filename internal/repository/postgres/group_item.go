package repopostgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
	"github.com/hokkung/release-management-service/pkg/gorem"
	"gorm.io/gorm"
)

type GroupItem struct {
	*gorem.BaseRepository[domain.GroupItem]
}

func NewGroupItem(db *gorm.DB) *GroupItem {
	return &GroupItem{
		BaseRepository: gorem.NewBaseRepository[domain.GroupItem](db),
	}
}

func (r *GroupItem) FindByCommitSHAs(ctx context.Context, shas []string) ([]domain.GroupItem, error) {
	return gorm.G[domain.GroupItem](r.GetDB(ctx)).Where("commit_sha IN ?", shas).Find(ctx)
}

func (r *GroupItem) FindByGroupID(ctx context.Context, groupID uuid.UUID) ([]domain.GroupItem, error) {
	return gorm.G[domain.GroupItem](r.GetDB(ctx)).Where("group_id = ?", groupID.String()).Find(ctx)
}

func (r *GroupItem) FindByGroupItemFilter(ctx context.Context, filter *domain.GroupItemFilter) ([]domain.GroupItem, error) {
	filters := make(map[string]any)
	if len(filter.GroupIDs) > 0 {
		filters["group_id"] = filter.GroupIDs
	}
	if len(filter.ReleasePlanIDs) > 0 {
		filters["release_plan_id"] = filter.ReleasePlanIDs
	}
	return r.BaseRepository.FindByFilter(ctx, filters)
}
