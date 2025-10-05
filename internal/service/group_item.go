package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
)

type GroupItem struct {
	groupItemRepository GroupItemRepository
}

func NewGroupItem(groupItemRepository GroupItemRepository) *GroupItem {
	return &GroupItem{
		groupItemRepository: groupItemRepository,
	}
}

type CreateGroupItemRequest struct {
	CommitSHA      string
	CommitAuthor   string
	CommitMesssage string
	ReleasePlanID  uuid.UUID
}

func (s *GroupItem) Create(ctx context.Context, req *CreateGroupItemRequest) error {
	ent := &domain.GroupItem{
		UIDModel: domain.UIDModel{
			ID: uuid.New(),
		},
		CommitSHA:      req.CommitSHA,
		CommitAuthor:   req.CommitAuthor,
		CommitMesssage: req.CommitMesssage,
		ReleasePlanID:  req.ReleasePlanID,
	}
	err := s.groupItemRepository.Create(ctx, ent)
	if err != nil {
		panic(err)
	}
	return nil
}
