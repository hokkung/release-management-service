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

type ListRequest struct {
	ReleasePlanIDs []uuid.UUID
}

type ListResponse struct {
	Entities []*ReleaseDto
}

type GroupDto struct {
	ID            uuid.UUID
	Name          string
	Status        string
	RepositoryID  uuid.UUID
	ReleasePlanID uuid.UUID
	GroupItems    []domain.GroupItem
}

type ReleaseDto struct {
	ReleaseID uuid.UUID
	Group     []GroupDto
	UnGroups  []domain.GroupItem
}

type CreateGroupItemRequest struct {
	CommitSHA      string
	CommitAuthor   string
	CommitMesssage string
	ReleasePlanID  uuid.UUID
}

type CreateIfNotExistRequest struct {
	Items []*CreateGroupItemRequest
}

type MoveRequest struct {
	FromGroupID uuid.UUID
	ToGroupID   uuid.UUID
	GroupItemID uuid.UUID
}
