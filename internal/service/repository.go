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
		fmt.Println("repository is up-to-date")
		return nil
	}
	releasePlan, err := s.createOrUpdateReleasePlan(ctx, ent, resp)
	if err != nil {
		return err
	}

	var commitsToBeCreated []*CreateGroupItemRequest
	for _, commit := range resp.Commits {
		// if merge strategy is "Create a merge commit"
		// skipping merge request commit created by Github
		if len(commit.Parents) > 1 {
			continue
		}
		fmt.Printf("commit: %+v \n", *commit.Commit.Message)
		commitsToBeCreated = append(commitsToBeCreated, &CreateGroupItemRequest{
			CommitSHA:      *commit.SHA,
			CommitAuthor:   *commit.Commit.Author.Email,
			CommitMesssage: *commit.Commit.Message,
			ReleasePlanID:  releasePlan.ID,
		})
	}
	groupItems, err := s.groupItemService.CreatesIfNotExist(ctx, &CreateIfNotExistRequest{
		Items: commitsToBeCreated,
	})
	fmt.Printf("groupItems: %+v \n", groupItems)

	return nil
}

func (s *Repository) createOrUpdateReleasePlan(
	ctx context.Context,
	ent *domain.Repository,
	resp *githuby.GetLatestCommitByBranchResponse,
) (*domain.ReleasePlan, error) {
	ongoingReleasePlans, err := s.releasePlanService.FindOngoingReleasePlans(ctx, &FindOngoingReleasePlansRequest{
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

	newReleasePlan, err := s.releasePlanService.Create(ctx, &CreateReleasePlanRequest{
		RepositoryID:           ent.ID,
		LatestTagCommit:        *resp.LatestTag.Commit.SHA,
		LatestMainBranchCommit: resp.HeadSHA,
	})
	if err != nil {
		return nil, err
	}
	return newReleasePlan, nil
}
