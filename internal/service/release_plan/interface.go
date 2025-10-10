package release_plan

import (
	"context"

	"github.com/hokkung/release-management-service/internal/service/group"
)

type GroupService interface {
	List(ctx context.Context, req *group.ListRequest) (*group.ListResponse, error)
}
