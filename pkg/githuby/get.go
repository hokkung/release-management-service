package githuby

import (
	"context"

	"github.com/google/go-github/v75/github"
)

type GetByRepositoryNameRequest struct {
	Name  string
	Owner string
}

type GetByRepositoryNameResponse struct {
	Repository *github.Repository
}

func (g *Github) GetByRepositoryName(ctx context.Context, req *GetByRepositoryNameRequest) (*GetByRepositoryNameResponse, error) {
	r, _, err := g.client.Repositories.Get(ctx, req.Owner, req.Name)
	if err != nil {
		return nil, err
	}
	return &GetByRepositoryNameResponse{
		Repository: r,
	}, nil
}
