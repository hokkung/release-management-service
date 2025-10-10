package domain

import (
	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/pkg/gorem"
)

type GroupStatus int

const (
	UnknownGroupStatus GroupStatus = iota
	TestingGroupStatus
	FailGroupStatus
	UatGroupStatus
	WaitingToDeploy
)

func (gs GroupStatus) String() string {
	switch gs {
	case TestingGroupStatus:
		return "TESTING"
	case FailGroupStatus:
		return "FAIL"
	case UatGroupStatus:
		return "UAT"
	case WaitingToDeploy:
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
		return WaitingToDeploy
	default:
		return UnknownGroupStatus
	}
}

type Group struct {
	UIDModel

	Name          string
	Status        string
	RepositoryID  uuid.UUID
	ReleasePlanID uuid.UUID
}

func (e *Group) TableName() string {
	return "rms.groups"
}

type GroupRepository interface {
	gorem.BaseRepositoryInt[Group]
}
