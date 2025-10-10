package model

import (
	"time"

	"github.com/google/uuid"
)

type ListReleasePlanRequest struct {
	*Pagination `json:"pagination"`
}

type ListReleasePlanResponse struct {
	ReleasePlans []ReleasePlan `json:"releasePlans"`
}

type ReleasePlan struct {
	ID                     uuid.UUID  `json:"id"`
	TargetDeployDate       *time.Time `json:"targetDeployDate"`
	Note                   *string    `json:"note"`
	LatestTagCommit        string     `json:"latestTag"`
	LatestMainBranchCommit string     `json:"latestMainBranchCommmit"`
	RepositoryID           uuid.UUID  `json:"repositoryId"`
	Groups                 []Group    `json:"groups"`
}

type SummaryReleasePlan struct {
}
