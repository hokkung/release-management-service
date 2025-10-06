package domain

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type ReleasePlanStatus string

var (
	WaitingToDeployReleasePlanStatus ReleasePlanStatus = "WAITING_TO_DEPLOY"
	TestingReleasePlanStatus         ReleasePlanStatus = "TESTING"
	FailReleasePlanStatus            ReleasePlanStatus = "FAIL"
	UatReleasePlanStatus             ReleasePlanStatus = "UAT"
)

type ReleasePlan struct {
	UIDModel

	fromCommit             string
	toCommit               string
	targetDeployDate       sql.NullTime
	Note                   sql.NullString
	LatestTagCommit        string
	LatestMainBranchCommit string
	RepositoryID           uuid.UUID
	Status                 string
}

func (e *ReleasePlan) TableName() string {
	return "rms.release_plans"
}

type ReleasePlanFilter struct {
	RepositoryIDs []uuid.UUID
}

type ReleasePlanRepository interface {
	Create(ctx context.Context, ent *ReleasePlan) error
	Save(ctx context.Context, ent *ReleasePlan) error
	FindByLatestMainBranchCommitAndNotInStatus(ctx context.Context, latestMainBranchCommit string, statuses []string) ([]ReleasePlan, error)
	FindByNotInStatus(ctx context.Context, statuses []string) ([]ReleasePlan, error)
	FindByFilter(ctx context.Context, filter *ReleasePlanFilter) ([]ReleasePlan, error)
}
