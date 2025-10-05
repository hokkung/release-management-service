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

func (s *ReleasePlan) Create(ctx context.Context, req *CreateReleasePlanRequest) error {
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
	return nil
}
