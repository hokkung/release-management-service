package group_item

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
)

type GroupItem struct {
	groupItemRepository domain.GroupItemRepository
}

func NewGroupItem(groupItemRepository domain.GroupItemRepository) *GroupItem {
	return &GroupItem{
		groupItemRepository: groupItemRepository,
	}
}

func (s *GroupItem) Create(ctx context.Context, req *CreateGroupItemRequest) (*domain.GroupItem, error) {
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
	return ent, nil
}

func (s *GroupItem) CreatesIfNotExist(ctx context.Context, req *CreateIfNotExistRequest) ([]*domain.GroupItem, error) {
	var shas []string
	shaToRequest := make(map[string]*CreateGroupItemRequest)
	for _, item := range req.Items {
		shas = append(shas, item.CommitSHA)
		shaToRequest[item.CommitSHA] = item
	}
	groupItems, err := s.groupItemRepository.FindByCommitSHAs(ctx, shas)
	if err != nil {
		return nil, err
	}
	for _, gi := range groupItems {
		delete(shaToRequest, gi.CommitSHA)
	}
	var itemsToBeCreated []*domain.GroupItem
	for _, s := range shaToRequest {
		itemsToBeCreated = append(itemsToBeCreated, &domain.GroupItem{
			UIDModel: domain.UIDModel{
				ID: uuid.New(),
			},
			CommitSHA:      s.CommitSHA,
			CommitAuthor:   s.CommitAuthor,
			CommitMesssage: s.CommitMesssage,
			ReleasePlanID:  s.ReleasePlanID,
		})
	}
	if err := s.Creates(ctx, itemsToBeCreated); err != nil {
		return nil, err
	}
	return itemsToBeCreated, nil
}

func (s *GroupItem) Creates(ctx context.Context, ents []*domain.GroupItem) error {
	return s.groupItemRepository.Creates(ctx, ents)
}

func (s *GroupItem) Move(ctx context.Context, req *MoveRequest) error {
	ent, exist, err := s.groupItemRepository.FindByKey(ctx, req.GroupItemID)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("entity not found")
	}
	if req.FromGroupID.String() != ent.GroupID.String() {
		return fmt.Errorf("group id is invalid")
	}
	ent.GroupID = &req.ToGroupID
	err = s.groupItemRepository.Save(ctx, ent)
	if err != nil {
		return err
	}
	return nil
}

func (s *GroupItem) UnassignByGroupID(ctx context.Context, groupID uuid.UUID) error {
	ents, err := s.groupItemRepository.FindByGroupID(ctx, groupID)
	if err != nil {
		return err
	}
	for _, ent := range ents {
		ent.GroupID = nil
		err = s.groupItemRepository.Save(ctx, &ent)
		if err != nil {
			return err
		}
	}
	return nil
}
