package domain

import "github.com/google/uuid"

type PullRequestGroupItem struct {
	UIDModel

	Status       string
	PullRequestID uuid.UUID
}
