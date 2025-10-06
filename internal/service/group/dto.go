package group

import (
	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
)

type CreateGroupRequest struct {
	Name          string
	RepositoryID  uuid.UUID
	ReleasePlanID uuid.UUID
}

type UpdateStatusRequest struct {
	GroupID uuid.UUID
	Status  string
}

type UpdateStatusResponse struct {
	Entity *domain.Group
}
