package group_item

import "github.com/google/uuid"

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
