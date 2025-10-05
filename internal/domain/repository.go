package domain

import (
	"database/sql"

	"github.com/google/uuid"
)

type RepositoryStatus string

var RegisteredRepositoryStatus RepositoryStatus = "REGISTERED"

type Repository struct {
	UIDModel

	Owner          string
	Name           string
	Url            string
	MainBranchName string
	Status         string
	LatestSyncAt   sql.NullTime
}

func (e *Repository) TableName() string {
	return "rms.repositories"
}

func NewRepository() *Repository {
	return &Repository{
		UIDModel: UIDModel{
			ID: uuid.New(),
		},
	}
}
