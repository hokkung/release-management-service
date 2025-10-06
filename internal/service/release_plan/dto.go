package release_plan

import (
	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
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
	Entities []domain.ReleasePlan
}
