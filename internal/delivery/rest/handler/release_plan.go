package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/hokkung/release-management-service/internal/delivery/rest/model"
	"github.com/hokkung/release-management-service/internal/service/release_plan"
)

type ReleasePlanService interface {
	List(ctx context.Context, req *release_plan.ListRequest) (*release_plan.ListResponse, error)
	ListSummary(ctx context.Context, req *release_plan.ListRequest) (*release_plan.ListResponse, error)
}

type ReleasePlanHandler struct {
	service ReleasePlanService
}

func NewReleasePlanHandler(service ReleasePlanService) *ReleasePlanHandler {
	return &ReleasePlanHandler{
		service: service,
	}
}

func (h *ReleasePlanHandler) List(c *fiber.Ctx) error {
	var req model.ListReleasePlanRequest
	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(err))
	}
	resp, err := h.service.List(c.UserContext(), &release_plan.ListRequest{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(handleAPIError(err))
	}

	releases := make([]model.ReleasePlan, 0, len(resp.Entities))
	for _, ent := range releases {
		releases = append(releases, model.ReleasePlan{

		})
	}
	return nil
}

func (h *ReleasePlanHandler) SetTargetDeployDate(c *fiber.Ctx) error {
	return nil
}
