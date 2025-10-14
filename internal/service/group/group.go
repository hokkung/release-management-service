package group

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
	"github.com/hokkung/release-management-service/pkg/gorem"
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
		UIDModel: gorem.UIDModel{
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
	err := s.repository.DeleteByID(ctx, groupID.String())
	if err != nil {
		return err
	}

	err = s.groupItemService.UnassignByGroupID(ctx, groupID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Group) ListByIDs(ctx context.Context, groupIds []uuid.UUID) ([]domain.Group, error) {
	if len(groupIds) == 0 {
		return []domain.Group{}, nil
	}

	groups, err := s.repository.FindByGroupFilter(ctx, &domain.GroupFilter{
		GroupIDs: groupIds,
	})
	if err != nil {
		return nil, err
	}

	return groups, nil
}


func (s *Group) ListByReleasePlanIDs(ctx context.Context, releasePlanIDs []uuid.UUID) ([]domain.Group, error) {
	if len(releasePlanIDs) == 0 {
		return []domain.Group{}, nil
	}

	groups, err := s.repository.FindByGroupFilter(ctx, &domain.GroupFilter{
		ReleasePlanIDs: releasePlanIDs,
	})
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (s *Group) GetLowestStatusLevel(ctx context.Context, groupIDs []uuid.UUID) (domain.GroupStatus, error) {
	groups, err := s.repository.FindByGroupFilter(ctx, &domain.GroupFilter{
		GroupIDs: groupIDs,
	})
	if err != nil {
		return domain.UnknownGroupStatus, err
	}

	if len(groups) == 0 {
		return domain.UnknownGroupStatus, fmt.Errorf("groups cannot be empty")
	}

	if len(groups) == 1 {
		return domain.NewGroupStatus(groups[0].Status), nil
	}

	currentStatus := domain.NewGroupStatus(groups[0].Status)
	for i := 1; i < len(groups); i++ {
		candidateStatus := domain.NewGroupStatus(groups[i].Status)
		if candidateStatus < currentStatus {
			currentStatus = candidateStatus
		}
	}

	return currentStatus, nil
}
