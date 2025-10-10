package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
	"github.com/hokkung/release-management-service/internal/service/group_item"
	"github.com/hokkung/release-management-service/internal/service/release_plan"
	"github.com/hokkung/release-management-service/pkg/githuby"
)

type Repository struct {
	repository         domain.RepositoryRepository
	groupItemService   GroupItemService
	githubService      GitHubService
	releasePlanService ReleasePlanService
}

func NewRepository(
	repository domain.RepositoryRepository,
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

func (s *Repository) Register(ctx context.Context, req *RegisterRequest) error {
	resp, err := s.githubService.GetByRepositoryName(ctx, &githuby.GetByRepositoryNameRequest{
		Owner: req.Owner,
		Name:  req.Name,
	})
	if err != nil {
		return err
	}

	err = s.Create(ctx, &CreateRequest{
		Name:  *resp.Repository.Name,
		Url:   *resp.Repository.URL,
		Owner: *resp.Repository.Owner.Login,
	})
	if err != nil {
		return err
	}
	return nil
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
		fmt.Println("repository is up-to-date")
		return nil
	}
	releasePlan, err := s.createOrUpdateReleasePlan(ctx, ent, resp)
	if err != nil {
		return err
	}

	var commitsToBeCreated []*group_item.CreateGroupItemRequest
	for _, commit := range resp.Commits {
		// if merge strategy is "Create a merge commit"
		// skipping merge request commit created by Github
		if NoParentSyncCommitType == req.SyncCommitType && len(commit.Parents) > 1 {
			continue
		}
		commitsToBeCreated = append(commitsToBeCreated, &group_item.CreateGroupItemRequest{
			CommitSHA:      *commit.SHA,
			CommitAuthor:   *commit.Commit.Author.Email,
			CommitMesssage: *commit.Commit.Message,
			ReleasePlanID:  releasePlan.ID,
		})
	}
	_, err = s.groupItemService.CreatesIfNotExist(ctx, &group_item.CreateIfNotExistRequest{
		Items: commitsToBeCreated,
	})
	ent.Status = string(domain.ActiveRepositoryStatus)
	err = s.repository.Save(ctx, ent)
	if err != nil {
		return err
	}
	return nil
}

func (s *Repository) createOrUpdateReleasePlan(
	ctx context.Context,
	ent *domain.Repository,
	resp *githuby.GetLatestCommitByBranchResponse,
) (*domain.ReleasePlan, error) {
	ongoingReleasePlans, err := s.releasePlanService.FindOngoingReleasePlans(ctx, &release_plan.FindOngoingReleasePlansRequest{
		LatestMainBranchCommit: resp.HeadSHA,
	})
	if err != nil {
		return nil, err
	}

	if len(ongoingReleasePlans.Entities) > 0 {
		ongoingReleasePlan := ongoingReleasePlans.Entities[0]
		if ongoingReleasePlan.LatestMainBranchCommit != resp.HeadSHA {
			ongoingReleasePlan.LatestMainBranchCommit = resp.HeadSHA
			err := s.releasePlanService.Update(ctx, &ongoingReleasePlan)
			if err != nil {
				return nil, err
			}
			return &ongoingReleasePlan, nil
		}
		return &ongoingReleasePlan, nil
	}

	newReleasePlan, err := s.releasePlanService.Create(ctx, &release_plan.CreateReleasePlanRequest{
		RepositoryID:           ent.ID,
		LatestTagCommit:        *resp.LatestTag.Commit.SHA,
		LatestMainBranchCommit: resp.HeadSHA,
	})
	if err != nil {
		return nil, err
	}
	return newReleasePlan, nil
}

func (s *Repository) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	ents, err := s.repository.FindActive(ctx)
	if err != nil {
		return nil, err
	}
	return &ListResponse{Entites: ents}, nil
}
