package githuby

import (
	"context"
	"fmt"

	"github.com/google/go-github/v75/github"
)

type GetLatestCommitByBranchRequest struct {
	BranchName     string
	Owner          string
	RepositoryName string
}

type GetLatestCommitByBranchResponse struct {
	HeadSHA   string
	AheadBy   int
	Commits   []*github.RepositoryCommit
	LatestTag *github.RepositoryTag
	IsDiff    bool
}

func (g *Github) HasUntaggedCommitOnMainBranch(ctx context.Context, req *GetLatestCommitByBranchRequest) (*GetLatestCommitByBranchResponse, error) {
	// 1. Get the latest tag
	latestTag, err := g.getLatestTag(ctx, req.Owner, req.RepositoryName)
	if err != nil {
		return nil, fmt.Errorf("get tags: %w", err)
	}
	if latestTag == nil {
		return nil, fmt.Errorf("no tags found in repo")
	}

	// 2. Get latest commit on main branch
	branch, _, err := g.client.Repositories.GetBranch(ctx, req.Owner, req.RepositoryName, req.BranchName, 0)
	if err != nil {
		return nil, fmt.Errorf("get main branch: %w", err)
	}
	headSHA := branch.GetCommit().GetSHA()

	// 3. Compare commits between latest tag and main branch
	comp, _, err := g.client.Repositories.CompareCommits(ctx, req.Owner, req.RepositoryName, *latestTag.Commit.SHA, headSHA, &github.ListOptions{
		PerPage: 50,
	})
	if err != nil {
		return nil, fmt.Errorf("compare commits: %w", err)
	}

	// If there are commits ahead of the tag, we have differences
	var isDiff bool = false
	if comp.GetAheadBy() > 0 {
		isDiff = true
	}

	return &GetLatestCommitByBranchResponse{
		LatestTag: latestTag,
		HeadSHA:   headSHA,
		AheadBy:   comp.GetAheadBy(),
		Commits:   comp.Commits,
		IsDiff:    isDiff,
	}, nil
}
