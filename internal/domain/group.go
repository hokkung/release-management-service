package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/pkg/gorem"
)

type GroupStatus int

const (
	UnknownGroupStatus GroupStatus = iota
	TestingGroupStatus
	FailGroupStatus
	UatGroupStatus
	WaitingToDeployGroupStatus
)

func (gs GroupStatus) String() string {
	switch gs {
	case TestingGroupStatus:
		return "TESTING"
	case FailGroupStatus:
		return "FAIL"
	case UatGroupStatus:
		return "UAT"
	case WaitingToDeployGroupStatus:
		return "WAITING_TO_DEPLOY"
	default:
		return "UNKNOWN"
	}
}

func NewGroupStatus(status string) GroupStatus {
	switch status {
	case "TESTING":
		return TestingGroupStatus
	case "FAIL":
		return FailGroupStatus
	case "UAT":
		return UatGroupStatus
	case "WAITING_TO_DEPLOY":
		return WaitingToDeployGroupStatus
	default:
		return UnknownGroupStatus
	}
}

type Group struct {
	gorem.UIDModel

	Name          string
	Status        string
	RepositoryID  uuid.UUID
	ReleasePlanID uuid.UUID
	GroupItems    []GroupItem `gorm:"-"`
}

func (e Group) TableName() string {
	return "rms.groups"
}

func (e Group) PrimaryKey() string {
	return "id"
}

type GroupFilter struct {
	GroupIDs       []uuid.UUID
	ReleasePlanIDs []uuid.UUID
}

type GroupRepository interface {
	gorem.Repository[Group]
	FindByGroupFilter(ctx context.Context, filter *GroupFilter) ([]Group, error)
}
