package domain

import (
	"time"

	"github.com/google/uuid"
)

type Release struct {
	UIDModel

	Name         string
	TargetDate   time.Time
	Note         *string
	RepositoryID uuid.UUID
	Status       string
}
