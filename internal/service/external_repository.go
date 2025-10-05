package service

import (
	"context"

	"github.com/hokkung/release-management-service/internal/domain"
)

type RepositoryRepository interface {
	Create(ctx context.Context, ent *domain.Repository) error
	FindByKey(ctx context.Context, key interface{}) (*domain.Repository, bool, error)
	FindByName(ctx context.Context, name string) (*domain.Repository, bool, error)
}

type GroupItemRepository interface {
	Create(ctx context.Context, ent *domain.GroupItem) error
	Creates(ctx context.Context, ents []*domain.GroupItem) error
	FindByCommitSHAs(ctx context.Context, shas []string) ([]domain.GroupItem, error)
}

type ReleasePlanRepository interface {
	Create(ctx context.Context, ent *domain.ReleasePlan) error
	Save(ctx context.Context, ent *domain.ReleasePlan) error
	FindByLatestMainBranchCommitAndNotInStatus(ctx context.Context, latestMainBranchCommit string, statuses []string) ([]domain.ReleasePlan, error)
	FindByNotInStatus(ctx context.Context, statuses []string) ([]domain.ReleasePlan, error)
}
