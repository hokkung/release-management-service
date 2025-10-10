package domain

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
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
	Create(ctx context.Context, ent *Repository) error
	FindByKey(ctx context.Context, key interface{}) (*Repository, bool, error)
	FindByName(ctx context.Context, name string) (*Repository, bool, error)
	FindActive(ctx context.Context) ([]Repository, error)
	Save(ctx context.Context, ent *Repository) error
}
