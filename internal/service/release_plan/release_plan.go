package release_plan

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
	"github.com/hokkung/release-management-service/internal/service/group"
	"github.com/hokkung/release-management-service/pkg/gorem"
)

type ReleasePlan struct {
	repository       domain.ReleasePlanRepository
	groupService     GroupService
	groupItemService GroupItemService
}

func NewReleasePlan(repository domain.ReleasePlanRepository, groupService GroupService, groupItemService GroupItemService) *ReleasePlan {
	return &ReleasePlan{
		repository:       repository,
		groupService:     groupService,
		groupItemService: groupItemService,
	}
}

func (s *ReleasePlan) Create(ctx context.Context, req *CreateReleasePlanRequest) (*domain.ReleasePlan, error) {
	ent := &domain.ReleasePlan{
		UIDModel: gorem.UIDModel{
			ID: uuid.New(),
		},
		LatestTagCommit:        req.LatestTagCommit,
		LatestMainBranchCommit: req.LatestMainBranchCommit,
		RepositoryID:           req.RepositoryID,
		Status:                 string(domain.TestingReleasePlanStatus),
	}
	err := s.repository.Create(ctx, ent)
	if err != nil {
		panic(err)
	}
	return ent, nil
}

func (s *ReleasePlan) FindOngoingReleasePlans(ctx context.Context, req *FindOngoingReleasePlansRequest) (*FindOngoingReleasePlansResponse, error) {
	ents, err := s.repository.FindByNotInStatus(ctx, []string{
		string(domain.WaitingToDeployReleasePlanStatus),
	})
	if err != nil {
		return nil, err
	}
	return &FindOngoingReleasePlansResponse{
		Entities: ents,
	}, nil
}

func (s *ReleasePlan) UpdateTargetDeployDateAndNote(ctx context.Context, req *UpdateTargetDeployDateAndNoteRequest) error {
	ent, exist, err := s.repository.FindByKey(ctx, req.ID)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("entity not found")
	}

	if req.TargetDeployDate != nil {
		ent.TargetDeployDate = sql.NullTime{Time: *req.TargetDeployDate, Valid: true}
	} else {
		ent.TargetDeployDate = sql.NullTime{}
	}

	if req.Note != nil {
		ent.Note = sql.NullString{String: *req.Note, Valid: true}
	} else {
		ent.Note = sql.NullString{}
	}

	return s.Update(ctx, ent)
}

func (s *ReleasePlan) Update(ctx context.Context, ent *domain.ReleasePlan) error {
	return s.repository.Save(ctx, ent)
}

func (s *ReleasePlan) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	ents, err := s.repository.FindByReleasePlanFilter(ctx, &domain.ReleasePlanFilter{
		RepositoryIDs: req.RepositoryIDs,
	})
	if err != nil {
		return nil, err
	}
	releasePlanDtos := make([]ReleasePlanDto, 0, len(ents))
	for _, ent := range ents {
		releasePlanDtos = append(releasePlanDtos, ReleasePlanDto{
			ID:                     ent.ID,
			TargetDeployDate:       &ent.TargetDeployDate.Time,
			Note:                   &ent.Note.String,
			LatestTagCommit:        ent.LatestTagCommit,
			LatestMainBranchCommit: ent.LatestMainBranchCommit,
			RepositoryID:           ent.RepositoryID,
			Status:                 ent.Status,
		})
	}
	return &ListResponse{
		Entities: releasePlanDtos,
	}, nil
}

func (s *ReleasePlan) ListSummary(ctx context.Context, req *ListSummaryRequest) (*ListSummaryResponse, error) {
	ents, err := s.repository.FindByReleasePlanFilter(ctx, &domain.ReleasePlanFilter{
		RepositoryIDs: req.RepositoryIDs,
	})
	if err != nil {
		return nil, err
	}

	releasePlanDtos := make([]*ReleasePlanDto, 0, len(ents))
	releaseIDs := make([]uuid.UUID, 0, len(ents))
	for _, ent := range ents {
		dto := &ReleasePlanDto{
			ID:                     ent.ID,
			LatestTagCommit:        ent.LatestTagCommit,
			LatestMainBranchCommit: ent.LatestMainBranchCommit,
			RepositoryID:           ent.RepositoryID,
			Status:                 ent.Status,
			Groups:                 []group.GroupDto{},
			UnGroupItems:           []domain.GroupItem{},
		}
		if ent.TargetDeployDate.Valid {
			dto.TargetDeployDate = &ent.TargetDeployDate.Time
		}
		if ent.Note.Valid {
			dto.Note = &ent.Note.String
		}

		releaseIDs = append(releaseIDs, ent.ID)
		releasePlanDtos = append(releasePlanDtos, dto)
	}

	groupItemsResp, err := s.groupItemService.ListByReleasePlanIDs(ctx, releaseIDs)
	if err != nil {
		return nil, err
	}

	groupIDs := make([]uuid.UUID, 0, len(groupItemsResp))
	groupIdToGroupItems := make(map[uuid.UUID][]domain.GroupItem)
	releaseIdToGroupItems := make(map[uuid.UUID][]domain.GroupItem)
	for _, groupItemResp := range groupItemsResp {
		// have group id
		if groupItemResp.GroupID != nil {
			groupIDs = append(groupIDs, *groupItemResp.GroupID)
			if groups, ok := groupIdToGroupItems[*groupItemResp.GroupID]; ok {
				groups = append(groups, groupItemResp)
				groupIdToGroupItems[*groupItemResp.GroupID] = groups
			} else {
				groupIdToGroupItems[*groupItemResp.GroupID] = []domain.GroupItem{groupItemResp}
			}
			continue
		}

		//  does not have group id
		if items, ok := releaseIdToGroupItems[groupItemResp.ReleasePlanID]; ok {
			items = append(items, groupItemResp)
			releaseIdToGroupItems[groupItemResp.ReleasePlanID] = items
		} else {
			releaseIdToGroupItems[groupItemResp.ReleasePlanID] = []domain.GroupItem{groupItemResp}
		}
	}

	groups, err := s.groupService.ListByReleasePlanIDs(ctx, releaseIDs)
	if err != nil {
		return nil, err
	}

	releasePlanIdToGroups := make(map[uuid.UUID][]domain.Group)
	for _, groupEnt := range groups {
		if groups, ok := releasePlanIdToGroups[groupEnt.ReleasePlanID]; ok {
			groups = append(groups, groupEnt)
			releasePlanIdToGroups[groupEnt.ReleasePlanID] = groups
		} else {
			releasePlanIdToGroups[groupEnt.ReleasePlanID] = []domain.Group{groupEnt}
		}
	}

	for _, releasePlanDto := range releasePlanDtos {
		var groupDtos []group.GroupDto
		if groups, ok := releasePlanIdToGroups[releasePlanDto.ID]; ok {
			for _, g := range groups {
				groupDto := group.GroupDto{
					ID:            g.ID,
					Name:          g.Name,
					Status:        g.Status,
					RepositoryID:  g.RepositoryID,
					ReleasePlanID: g.ReleasePlanID,
				}
				if items, ok := groupIdToGroupItems[groupDto.ID]; ok {
					groupDto.GroupItems = items
				}
				groupDtos = append(groupDtos, groupDto)
			}
		}
		releasePlanDto.Groups = groupDtos
		if items, ok := releaseIdToGroupItems[releasePlanDto.ID]; ok {
			releasePlanDto.UnGroupItems = items
		}
	}

	return &ListSummaryResponse{
		Entities: releasePlanDtos,
	}, nil
}

func (s *ReleasePlan) HandleGroupStatusUpdated(ctx context.Context, req *UpdateStatusRequest) error {
	ent, exist, err := s.repository.FindByKey(ctx, req.ReleasePlanID)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("entity not found")
	}

	items, err := s.groupItemService.ListByReleasePlanIDs(ctx, []uuid.UUID{ent.ID})
	if err != nil {
		return err
	}

	groupIDs := make([]uuid.UUID, 0, len(items))
	for _, item := range items {
		if item.GroupID == nil {
			ent.Status = string(domain.TestingReleasePlanStatus)
			return s.Update(ctx, ent)
		}

		groupIDs = append(groupIDs, *item.GroupID)
	}

	status, err := s.groupService.GetLowestStatusLevel(ctx, groupIDs)
	if err != nil {
		return err
	}
	releasePlanStatus := s.GetStatusByGroupStatus(status)
	ent.Status = string(releasePlanStatus)

	return s.Update(ctx, ent)
}

func (s *ReleasePlan) GetStatusByGroupStatus(groupStatus domain.GroupStatus) domain.ReleasePlanStatus {
	switch groupStatus {
	case domain.FailGroupStatus:
		return domain.FailReleasePlanStatus
	case domain.UatGroupStatus:
		return domain.UatReleasePlanStatus
	case domain.WaitingToDeployGroupStatus:
		return domain.WaitingToDeployReleasePlanStatus
	default:
		return domain.TestingReleasePlanStatus
	}
}
