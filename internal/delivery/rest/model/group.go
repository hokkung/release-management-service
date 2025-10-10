package model

import (
	"github.com/google/uuid"
)

type CreateGroupRequest struct {
	Name          string    `json:"name" required:"true"`
	RepositoryID  uuid.UUID `json:"repositoryId" required:"true"`
	ReleasePlanID uuid.UUID `json:"releasePlanId" required:"true"`
}

type CreateGroupResponse struct {
}

type MoveRequest struct {
	ToGroupID uuid.UUID `json:"toGroupId" required:"true"`
}

type MoveResponse struct {
}

type UpdateGroupStatusRequest struct {
	Status string `json:"status" required:"true"`
}

type UpdateGroupStatusResponse struct {
}

type RemoveRequest struct {
}

type RemoveResponse struct {
}

type APIGroupDataResponse APIDataResponse[Group]
type Group struct {
	ID         uuid.UUID   `json:"id"`
	Name       string      `json:"name"`
	Status     string      `json:"status"`
	GroupItems []GroupItem `json:"groupItems"`
}

type GroupStatusUpdated struct {
	GroupID uuid.UUID
}
