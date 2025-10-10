package release_plan

import (
	"time"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
	"github.com/hokkung/release-management-service/internal/service/group"
)

type CreateReleasePlanRequest struct {
	RepositoryID           uuid.UUID
	LatestTagCommit        string
	LatestMainBranchCommit string
}

type FindOngoingReleasePlansRequest struct {
	LatestMainBranchCommit string
}

type FindOngoingReleasePlansResponse struct {
	Entities []domain.ReleasePlan
}

type ListRequest struct {
	RepositoryIDs []uuid.UUID
}

type ListResponse struct {
	Entities []ReleasePlanDto
}

type ListSummaryRequest struct {
	RepositoryIDs []uuid.UUID
}

type ListSummaryResponse struct {
	Entities []*ReleasePlanDto
}

type ReleasePlanDto struct {
	ID                     uuid.UUID
	TargetDeployDate       *time.Time
	Note                   *string
	LatestTagCommit        string
	LatestMainBranchCommit string
	RepositoryID           uuid.UUID
	Status                 string
	Groups                 []group.GroupDto
	UnGroupItems           []domain.GroupItem
}

type UpdateTargetDeployDateAndNoteRequest struct {
	ID               uuid.UUID
	TargetDeployDate *time.Time
	Note             *string
}

type UpdateStatusRequest struct {
	GroupID uuid.UUID
	ReleasePlanID uuid.UUID
	RepositoryID uuid.UUID
}
