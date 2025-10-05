package githuby

import (
	"context"

	"github.com/google/go-github/v75/github"
)

type Github struct {
	client *github.Client
}

func New(client *github.Client) *Github {
	return &Github{
		client: client,
	}
}

func (g *Github) GetCurrentUser(ctx context.Context) (*github.User, error) {
	u, _, err := g.client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	return u, nil
}
