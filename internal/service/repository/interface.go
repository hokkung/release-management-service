package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
	"github.com/hokkung/release-management-service/internal/service/group"
	"github.com/hokkung/release-management-service/internal/service/release_plan"
	"github.com/hokkung/release-management-service/pkg/githuby"
)

type ReleasePlanService interface {
	Create(ctx context.Context, req *release_plan.CreateReleasePlanRequest) (*domain.ReleasePlan, error)
	Update(ctx context.Context, ent *domain.ReleasePlan) error
	FindOngoingReleasePlans(ctx context.Context, req *release_plan.FindOngoingReleasePlansRequest) (*release_plan.FindOngoingReleasePlansResponse, error)
}

type GitHubService interface {
	GetByRepositoryName(ctx context.Context, req *githuby.GetByRepositoryNameRequest) (*githuby.GetByRepositoryNameResponse, error)
	HasUntaggedCommitOnMainBranch(ctx context.Context, req *githuby.GetLatestCommitByBranchRequest) (*githuby.GetLatestCommitByBranchResponse, error)
}

type GroupItemService interface {
	Create(ctx context.Context, req *group.CreateGroupItemRequest) (*domain.GroupItem, error)
	Creates(ctx context.Context, ents []*domain.GroupItem) error
	CreatesIfNotExist(ctx context.Context, req *group.CreateIfNotExistRequest) ([]*domain.GroupItem, error)
	UnassignByGroupID(ctx context.Context, groupID uuid.UUID) error
}
