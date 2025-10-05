package service

import (
	"context"

	"github.com/hokkung/release-management-service/internal/domain"
	"github.com/hokkung/release-management-service/pkg/githuby"
)

type GitHubService interface {
	GetByRepositoryName(ctx context.Context, req *githuby.GetByRepositoryNameRequest) (*githuby.GetByRepositoryNameResponse, error)
	HasUntaggedCommitOnMainBranch(ctx context.Context, req *githuby.GetLatestCommitByBranchRequest) (*githuby.GetLatestCommitByBranchResponse, error)
}

type GroupItemService interface {
	Create(ctx context.Context, req *CreateGroupItemRequest) (*domain.GroupItem, error)
}

type ReleasePlanService interface {
	Create(ctx context.Context, req *CreateReleasePlanRequest) (*domain.ReleasePlan, error)
}
