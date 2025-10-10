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
	gorem.UIDModel

	Owner          string
	Name           string
	Url            string
	MainBranchName string
	Status         string
	LatestSyncAt   sql.NullTime
}

func (e Repository) TableName() string {
	return "rms.repositories"
}

func (e Repository) PrimaryKey() string {
	return "id"
}

func NewRepository() *Repository {
	return &Repository{
		UIDModel: gorem.UIDModel{
			ID: uuid.New(),
		},
	}
}

type RepositoryRepository interface {
	gorem.Repository[Repository]
}
