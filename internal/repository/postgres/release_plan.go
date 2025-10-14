package repopostgres

import (
	"context"

	"github.com/hokkung/release-management-service/internal/domain"
	"github.com/hokkung/release-management-service/pkg/gorem"
	"gorm.io/gorm"
)

type ReleasePlan struct {
	*gorem.BaseRepository[domain.ReleasePlan]
}

func NewReleasePlan(db *gorm.DB) *ReleasePlan {
	return &ReleasePlan{
		BaseRepository: gorem.NewBaseRepository[domain.ReleasePlan](db),
	}
}

func (r *ReleasePlan) FindByLatestMainBranchCommitAndNotInStatus(ctx context.Context, latestMainBranchCommit string, statuses []string) ([]domain.ReleasePlan, error) {
	ents, err := gorm.G[domain.ReleasePlan](r.GetDB(ctx)).Where("latest_main_branch_commit = ? AND status NOT IN ?", latestMainBranchCommit, statuses).Find(ctx)
	return ents, err
}

func (r *ReleasePlan) FindByNotInStatus(ctx context.Context, statuses []string) ([]domain.ReleasePlan, error) {
	ents, err := gorm.G[domain.ReleasePlan](r.GetDB(ctx)).Where("status NOT IN ?", statuses).Find(ctx)
	return ents, err
}

func (r *ReleasePlan) FindByReleasePlanFilter(ctx context.Context, filter *domain.ReleasePlanFilter) ([]domain.ReleasePlan, error) {
	filters := make(map[string]any)
	if len(filter.RepositoryIDs) > 0 {
		filters["repository_id"] = filter.RepositoryIDs
	}
	return r.BaseRepository.FindByFilter(ctx, filters)
}
