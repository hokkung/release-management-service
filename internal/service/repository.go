package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
)

type Repository struct {
	repository RepositoryRepository
}

func NewRepository(repository RepositoryRepository) *Repository {
	return &Repository{
		repository: repository,
	}
}

func (s *Repository) Create() error {
	ctx := context.Background()
	ent := &domain.Repository{
		UIDModel: domain.UIDModel{
			ID: uuid.New(),
		},
		Name:      "test",
		Url:       "localhost:8080",
		ServiceID: uuid.New(),
	}
	err := s.repository.Create(ctx, ent)
	if err != nil {
		panic(err)
	}

	fmt.Printf("ent: %+v", *ent)
	return nil
}
