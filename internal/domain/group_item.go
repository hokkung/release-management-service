package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/pkg/gorem"
)

type GroupItem struct {
	gorem.UIDModel

	CommitSHA      string
	CommitAuthor   string
	CommitMesssage string
	GroupID        *uuid.UUID
	ReleasePlanID  uuid.UUID
}

func (e GroupItem) TableName() string {
	return "rms.group_items"
}

func (e GroupItem) PrimaryKey() string {
	return "id"
}

type GroupItemFilter struct {
	GroupIDs       []uuid.UUID
	ReleasePlanIDs []uuid.UUID
}

type GroupItemRepository interface {
	gorem.Repository[GroupItem]
	FindByCommitSHAs(ctx context.Context, shas []string) ([]GroupItem, error)
	FindByGroupID(ctx context.Context, groupID uuid.UUID) ([]GroupItem, error)
	FindByGroupItemFilter(ctx context.Context, filter *GroupItemFilter) ([]GroupItem, error)
}
