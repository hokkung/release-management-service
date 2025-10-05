package repopostgres

import (
	"context"

	"github.com/hokkung/release-management-service/internal/domain"
	"github.com/hokkung/release-management-service/pkg/gorem"
	"gorm.io/gorm"
)

type GroupItem struct {
	gorem.BaseRepository[domain.GroupItem]
}

func NewGroupItem(db *gorm.DB) *GroupItem {
	return &GroupItem{
		BaseRepository: gorem.BaseRepository[domain.GroupItem](*gorem.NewBaseRepository[domain.GroupItem](db)),
	}
}

func (r *GroupItem) FindByCommitSHAs(ctx context.Context, shas []string) ([]domain.GroupItem, error) {
	return gorm.G[domain.GroupItem](r.GetDB(ctx)).Where("commit_sha IN ?", shas).Find(ctx)
}
