package service

import (
	"context"

	"github.com/hokkung/release-management-service/internal/domain"
)

type RepositoryRepository interface {
	Create(ctx context.Context, ent *domain.Repository) error
}
