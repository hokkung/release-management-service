package release_plan

import (
	"context"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
)

type GroupService interface {
	ListByIDs(ctx context.Context, groupIds []uuid.UUID) ([]domain.Group, error)
	ListByReleasePlanIDs(ctx context.Context, releasePlanIDs []uuid.UUID) ([]domain.Group, error)
	GetLowestStatusLevel(ctx context.Context, groupIDs []uuid.UUID) (domain.GroupStatus, error)
}

type GroupItemService interface {
	ListByReleasePlanIDs(ctx context.Context, releasePlanIDs []uuid.UUID) ([]domain.GroupItem, error)
}
