package model

import "github.com/google/uuid"

type GroupItem struct {
	ID             uuid.UUID `json:"id"`
	CommitSHA      string    `json:"commitSHA"`
	CommitAuthor   string    `json:"author"`
	CommitMesssage string    `json:"commitMessage"`
}
