package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
)

type ReleasePlan struct {
	repository ReleasePlanRepository
}

func NewReleasePlan(repository ReleasePlanRepository) *ReleasePlan {
	return &ReleasePlan{
		repository: repository,
	}
}

type CreateReleasePlanRequest struct {
	RepositoryID           uuid.UUID
	LatestTagCommit        string
	LatestMainBranchCommit string
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

type FindOngoingReleasePlansRequest struct {
	LatestMainBranchCommit string
}

type FindOngoingReleasePlansResponse struct {
	Entities []domain.ReleasePlan
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
