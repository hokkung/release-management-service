package group

import (
	"context"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
	"github.com/hokkung/release-management-service/internal/service/group_item"
)

type GroupItemService interface {
	Create(ctx context.Context, req *group_item.CreateGroupItemRequest) (*domain.GroupItem, error)
	Creates(ctx context.Context, ents []*domain.GroupItem) error
	CreatesIfNotExist(ctx context.Context, req *group_item.CreateIfNotExistRequest) ([]*domain.GroupItem, error)
	UnassignByGroupID(ctx context.Context, groupID uuid.UUID) error
}
