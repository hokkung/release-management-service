package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/pkg/gorem"
)

type GroupItem struct {
	UIDModel

	CommitSHA      string
	CommitAuthor   string
	CommitMesssage string
	GroupID        *uuid.UUID
	ReleasePlanID  uuid.UUID
}

func (e *GroupItem) TableName() string {
	return "rms.group_items"
}

type GroupItemRepository interface {
	gorem.BaseRepositoryInt[GroupItem]
	FindByCommitSHAs(ctx context.Context, shas []string) ([]GroupItem, error)
	FindByGroupID(ctx context.Context, groupID uuid.UUID) ([]GroupItem, error)
}
