package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
	"github.com/hokkung/release-management-service/pkg/githuby"
)

type Repository struct {
	repository         RepositoryRepository
	groupItemService   GroupItemService
	githubService      GitHubService
	releasePlanService ReleasePlanService
}

func NewRepository(
	repository RepositoryRepository,
	githubService GitHubService,
	groupItemService GroupItemService,
	releasePlanService ReleasePlanService,
) *Repository {
	return &Repository{
		repository:         repository,
		githubService:      githubService,
		groupItemService:   groupItemService,
		releasePlanService: releasePlanService,
	}
}

type CreateRequest struct {
	Name  string
	Url   string
	Owner string
}

func (s *Repository) Create(ctx context.Context, req *CreateRequest) error {
	ent := &domain.Repository{
		UIDModel: domain.UIDModel{
			ID: uuid.New(),
		},
		Name:           req.Name,
		Url:            req.Url,
		Owner:          req.Owner,
		MainBranchName: "main",
		Status:         string(domain.RegisteredRepositoryStatus),
	}
	err := s.repository.Create(ctx, ent)
	if err != nil {
		panic(err)
	}
	return nil
}

type RegisterRequest struct {
	Name string
}

func (s *Repository) Register(ctx context.Context, req *RegisterRequest) error {
	resp, err := s.githubService.GetByRepositoryName(ctx, &githuby.GetByRepositoryNameRequest{
		Owner: "hokkung",
		Name:  req.Name,
	})
	if err != nil {
		return err
	}

	err = s.Create(ctx, &CreateRequest{
		Name:  *resp.Repository.Name,
		Url:   *resp.Repository.URL,
		Owner: "hokkung",
	})
	if err != nil {
		return err
	}
	return nil
}

type SyncRequest struct {
	RepositoryName string
}

func (s *Repository) Sync(ctx context.Context, req *SyncRequest) error {
	ent, exist, err := s.repository.FindByName(ctx, req.RepositoryName)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("not found repository name: %s", req.RepositoryName)
	}
	resp, err := s.githubService.HasUntaggedCommitOnMainBranch(ctx, &githuby.GetLatestCommitByBranchRequest{
		BranchName:     ent.MainBranchName,
		Owner:          ent.Owner,
		RepositoryName: ent.Name,
	})
	if err != nil {
		return err
	}
	if !resp.IsDiff {
		return nil
	}
	releasePlan, err := s.releasePlanService.Create(ctx, &CreateReleasePlanRequest{
		RepositoryID:           ent.ID,
		LatestTagCommit:        *resp.LatestTag.Commit.SHA,
		LatestMainBranchCommit: resp.HeadSHA,
	})
	if err != nil {
		return err
	}
	var groupItems []*domain.GroupItem
	for _, commit := range resp.Commits {
		groupItem, err := s.groupItemService.Create(ctx, &CreateGroupItemRequest{
			CommitSHA:      *commit.Commit.SHA,
			CommitAuthor:   *commit.Commit.Author.Login,
			CommitMesssage: *commit.Commit.Message,
			ReleasePlanID:  releasePlan.ID,
		})
		if err != nil {
			return err
		}
		groupItems = append(groupItems, groupItem)
	}

	return nil
}
