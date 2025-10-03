package domain

import "github.com/google/uuid"

type PullRequeset struct {
	UIDModel

	Title        string
	Status       string
	RepositoryID uuid.UUID
}
