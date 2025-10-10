package model

import "github.com/google/uuid"

type CreateGroupRequest struct {
}

type CreateGroupResponse struct {
}

type MoveRequest struct {
}

type MoveResponse struct {
}

type UpdateStatusRequest struct {
}

type UpdateStatusResponse struct {
}

type RemoveRequest struct {
}

type RemoveResponse struct {
}

type Group struct {
	ID         uuid.UUID   `json:"id"`
	Name       string      `json:"name"`
	Status     string      `json:"status"`
	GroupItems []GroupItem `json:"groupItems"`
}
