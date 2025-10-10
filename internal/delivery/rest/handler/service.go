package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/domain"
	"github.com/hokkung/release-management-service/internal/service/group"
	"github.com/hokkung/release-management-service/internal/service/release_plan"
	"github.com/hokkung/release-management-service/internal/service/repository"
)

type GroupItemService interface {
	Move(ctx context.Context, req *group.MoveRequest) error
}

type GroupService interface {
	ListByIDs(ctx context.Context, groupIds []uuid.UUID) ([]domain.Group, error)
	Remove(ctx context.Context, groupID uuid.UUID) error
	Create(ctx context.Context, req *group.CreateGroupRequest) (*domain.Group, error)
	UpdateStatus(ctx context.Context, req *group.UpdateStatusRequest) (*group.UpdateStatusResponse, error)
}

type ReleasePlanService interface {
	List(ctx context.Context, req *release_plan.ListRequest) (*release_plan.ListResponse, error)
	ListSummary(ctx context.Context, req *release_plan.ListSummaryRequest) (*release_plan.ListSummaryResponse, error)
	UpdateTargetDeployDateAndNote(ctx context.Context, req *release_plan.UpdateTargetDeployDateAndNoteRequest) error
	HandleGroupStatusUpdated(ctx context.Context, req *release_plan.UpdateStatusRequest) error
}

type RepositoryService interface {
	Create(ctx context.Context, req *repository.CreateRequest) error
	Register(ctx context.Context, req *repository.RegisterRequest) error
	Sync(ctx context.Context, req *repository.SyncRequest) error
	List(ctx context.Context, req *repository.ListRequest) (*repository.ListResponse, error)
}
