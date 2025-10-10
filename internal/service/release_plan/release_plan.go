package release_plan

import (
	"context"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
	"github.com/hokkung/release-management-service/internal/service/group"
)

type ReleasePlan struct {
	repository   domain.ReleasePlanRepository
	groupService GroupService
}

func NewReleasePlan(repository domain.ReleasePlanRepository, groupService GroupService) *ReleasePlan {
	return &ReleasePlan{
		repository:   repository,
		groupService: groupService,
	}
}

func (s *ReleasePlan) Create(ctx context.Context, req *CreateReleasePlanRequest) (*domain.ReleasePlan, error) {
	ent := &domain.ReleasePlan{
		UIDModel: domain.UIDModel{
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
	return &ListResponse{
		Entities: ents,
	}, nil
}

func (s *ReleasePlan) ListSummary(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	ents, err := s.repository.FindByReleasePlanFilter(ctx, &domain.ReleasePlanFilter{
		RepositoryIDs: req.RepositoryIDs,
	})
	if err != nil {
		return nil, err
	}
	releasePlanDtos := make([]*ReleasePlanDto, 0, len(ents))
	releaseIDs := make([]uuid.UUID, 0, len(ents))
	for _, ent := range ents {
		releaseIDs = append(releaseIDs, ent.ID)
		releasePlanDtos = append(releasePlanDtos, &ReleasePlanDto{
			ID:                     ent.ID,
			TargetDeployDate:       &ent.TargetDeployDate.Time,
			Note:                   &ent.Note.String,
			LatestTagCommit:        ent.LatestTagCommit,
			LatestMainBranchCommit: ent.LatestMainBranchCommit,
			RepositoryID:           ent.RepositoryID,
			Status:                 ent.Status,
		})
	}
	groupsResp, err := s.groupService.List(ctx, &group.ListRequest{
		ReleasePlanIDs: releaseIDs,
	})
	if err != nil {
		return nil, err
	}
	releasePlanIDToGroups := make(map[uuid.UUID][]group.GroupDto)
	for _, g := range groupsResp.Entities {
		groups, ok := releasePlanIDToGroups[g.ReleasePlanID]
		if ok {
			groups = append(groups, *g)
			releasePlanIDToGroups[g.ReleasePlanID] = groups
		} else {
			releasePlanIDToGroups[g.ReleasePlanID] = []group.GroupDto{}
		}
	}
	for _, releasePlanDto := range releasePlanDtos {
		if groups, ok := releasePlanIDToGroups[releasePlanDto.ID]; ok {
			releasePlanDto.Groups = groups
		}
	}
	return &ListResponse{
		Entities: releasePlanDtos,
	}, nil
}
