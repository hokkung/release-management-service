package group

import (
	"context"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
)

type GroupService interface {
}

type GroupItemService interface {
	Create(ctx context.Context, req *CreateGroupItemRequest) (*domain.GroupItem, error)
	Creates(ctx context.Context, ents []*domain.GroupItem) error
	CreatesIfNotExist(ctx context.Context, req *CreateIfNotExistRequest) ([]*domain.GroupItem, error)
	UnassignByGroupID(ctx context.Context, groupID uuid.UUID) error
	ListByGroupIDs(ctx context.Context, groupIDs []uuid.UUID) ([]domain.GroupItem, error)
	ListByReleasePlanIDs(ctx context.Context, releasePlanIDs []uuid.UUID) ([]domain.GroupItem, error)
}
