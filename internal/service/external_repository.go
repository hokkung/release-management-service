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
}


type ReleasePlanRepository interface {
	Create(ctx context.Context, ent *domain.ReleasePlan) error
}
