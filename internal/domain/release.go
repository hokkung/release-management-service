package domain

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/pkg/gorem"
)

type ReleasePlanStatus string

var (
	TestingReleasePlanStatus         ReleasePlanStatus = "TESTING"
	FailReleasePlanStatus            ReleasePlanStatus = "FAIL"
	UatReleasePlanStatus             ReleasePlanStatus = "UAT"
	WaitingToDeployReleasePlanStatus ReleasePlanStatus = "WAITING_TO_DEPLOY"
)

type ReleasePlan struct {
	gorem.UIDModel

	TargetDeployDate       sql.NullTime
	Note                   sql.NullString
	LatestTagCommit        string
	LatestMainBranchCommit string
	RepositoryID           uuid.UUID
	Status                 string
}

func (e ReleasePlan) TableName() string {
	return "rms.release_plans"
}

func (e ReleasePlan) PrimaryKey() string {
	return "id"
}

type ReleasePlanFilter struct {
	RepositoryIDs []uuid.UUID
}

type ReleasePlanRepository interface {
	gorem.Repository[ReleasePlan]
	FindByLatestMainBranchCommitAndNotInStatus(ctx context.Context, latestMainBranchCommit string, statuses []string) ([]ReleasePlan, error)
	FindByNotInStatus(ctx context.Context, statuses []string) ([]ReleasePlan, error)
	FindByReleasePlanFilter(ctx context.Context, filter *ReleasePlanFilter) ([]ReleasePlan, error)
}
