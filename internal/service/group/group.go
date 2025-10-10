package group

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
)

type Group struct {
	repository       domain.GroupRepository
	groupItemService GroupItemService
}

func NewGroup(repository domain.GroupRepository, groupItemService GroupItemService) *Group {
	return &Group{
		repository:       repository,
		groupItemService: groupItemService,
	}
}

func (s *Group) Create(ctx context.Context, req *CreateGroupRequest) (*domain.Group, error) {
	ent := domain.Group{
		UIDModel: domain.UIDModel{
			ID: uuid.New(),
		},
		Name:          req.Name,
		Status:        domain.TestingGroupStatus.String(),
		RepositoryID:  req.RepositoryID,
		ReleasePlanID: req.ReleasePlanID,
	}
	err := s.repository.Create(ctx, &ent)
	if err != nil {
		return nil, err
	}
	return &ent, nil
}

func (s *Group) UpdateStatus(ctx context.Context, req *UpdateStatusRequest) (*UpdateStatusResponse, error) {
	ent, exist, err := s.repository.FindByKey(ctx, req.GroupID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("entity not found")
	}
	status := domain.NewGroupStatus(req.Status)
	if domain.UnknownGroupStatus == status {
		return nil, fmt.Errorf("invalid status")
	}
	ent.Status = status.String()
	err = s.repository.Save(ctx, ent)
	if err != nil {
		return nil, err
	}
	return &UpdateStatusResponse{
		Entity: ent,
	}, nil
}

func (s *Group) Remove(ctx context.Context, groupID uuid.UUID) error {
	err := s.groupItemService.UnassignByGroupID(ctx, groupID)
	if err != nil {
		return err
	}
	err = s.repository.DeleteByID(ctx, groupID.String())
	if err != nil {
		return err
	}
	return nil
}

func (s *Group) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	groups, err := s.repository.FindByGroupFilter(ctx, &domain.GroupFilter{
		ReleasePlanIDs: req.ReleasePlanIDs,
	})
	if err != nil {
		return nil, err
	}
	groupDtos := make([]*GroupDto, 0, len(groups))
	groupIDs := make([]uuid.UUID, 0, len(groups))
	for _, group := range groups {
		groupIDs = append(groupIDs, group.ID)
		groupDtos = append(groupDtos, &GroupDto{
			ID:            group.ID,
			Name:          group.Name,
			Status:        group.Status,
			RepositoryID:  group.RepositoryID,
			ReleasePlanID: group.ReleasePlanID,
		})
	}
	items, err := s.groupItemService.ListByGroupIDs(ctx, groupIDs)
	if err != nil {
		return nil, err
	}
	groupIdToItems := make(map[uuid.UUID][]domain.GroupItem)
	for _, item := range items {
		items, ok := groupIdToItems[*item.GroupID]
		if ok {
			items = append(items, item)
			groupIdToItems[*item.GroupID] = items
		} else {
			groupIdToItems[*item.GroupID] = []domain.GroupItem{}
		}
	}
	for _, groupDto := range groupDtos {
		if items, ok := groupIdToItems[groupDto.ID]; ok {
			groupDto.GroupItems = items
		}
	}
	return &ListResponse{
		Entities: groupDtos,
	}, nil
}
