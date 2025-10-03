package domain

import "github.com/google/uuid"

type PullRequest struct {
	UIDModel

	Status       string
	RepositoryID uuid.UUID
}
