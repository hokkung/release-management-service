package domain

import "github.com/google/uuid"

type GroupItem struct {
	UIDModel

	CommitSHA      string
	CommitAuthor   string
	CommitMesssage string
	GroupID        *uuid.UUID
	ReleasePlanID  uuid.UUID
}
