package domain

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/pkg/gorem"
)

type RepositoryStatus string

const (
	RegisteredRepositoryStatus RepositoryStatus = "REGISTERED"
	ActiveRepositoryStatus     RepositoryStatus = "ACTIVE"
)

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

type RepositoryRepository interface {
	gorem.BaseRepositoryInt[Repository]
}
