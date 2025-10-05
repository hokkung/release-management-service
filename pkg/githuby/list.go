package githuby

import (
	"context"

	"github.com/google/go-github/v75/github"
)

type ListRepositoryRequest struct {
	User string
}

type ListRepositoryResponse struct {
	Repositories []*github.Repository
}

func (g *Github) ListRepository(ctx context.Context, req *ListRepositoryRequest) (*ListRepositoryResponse, error) {
	opt := &github.RepositoryListByUserOptions{
		ListOptions: github.ListOptions{
			PerPage: 50,
		},
	}
	var repos []*github.Repository
	for {
		rs, resp, err := g.client.Repositories.ListByUser(ctx, req.User, opt)
		if err != nil {
			return nil, err
		}
		repos = append(repos, rs...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return &ListRepositoryResponse{
		Repositories: repos,
	}, nil
}

