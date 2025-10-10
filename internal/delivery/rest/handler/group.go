package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/delivery/rest/model"
	"github.com/hokkung/release-management-service/internal/service/group"
	"github.com/hokkung/release-management-service/internal/service/release_plan"
)

type Group struct {
	groupService       GroupService
	releasePlanService ReleasePlanService
}

func NewGroup(groupService GroupService, releasePlanService ReleasePlanService) *Group {
	return &Group{
		groupService:       groupService,
		releasePlanService: releasePlanService,
	}
}

// @Summary Create group API
// @Tags Group
// @Accept json
// @Produce json
// @Param request body model.CreateGroupRequest true "Group request"
// @Success 200 {object} model.APIGroupDataResponse
// @Router /api/v1/groups [post]
func (h *Group) CreateGroup(c *fiber.Ctx) error {
	var req model.CreateGroupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(err))
	}

	ent, err := h.groupService.Create(c.UserContext(), &group.CreateGroupRequest{
		Name:          req.Name,
		RepositoryID:  req.RepositoryID,
		ReleasePlanID: req.ReleasePlanID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(handleAPIError(err))
	}

	return c.JSON(model.APIDataResponse[model.Group]{
		Data: model.Group{
			ID:     ent.ID,
			Name:   ent.Name,
			Status: ent.Status,
		},
	})
}

// @Summary Update group status API
// @Tags Group
// @Accept json
// @Produce json
// @Param group_id path uuid.UUID true "Group ID"
// @Param request body model.UpdateGroupStatusRequest true "UpdateGroupStatus request"
// @Success 200 {object} model.APIGroupDataResponse
// @Router /api/v1/groups/{group_id}/update-status [post]
func (h *Group) UpdateStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(fmt.Errorf("id is invalid")))
	}
	groupID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(err))
	}

	var req model.UpdateGroupStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(err))
	}

	resp, err := h.groupService.UpdateStatus(c.UserContext(), &group.UpdateStatusRequest{
		GroupID: groupID,
		Status:  req.Status,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(handleAPIError(err))
	}

	err = h.releasePlanService.HandleGroupStatusUpdated(c.UserContext(), &release_plan.UpdateStatusRequest{
		GroupID:       resp.Entity.ID,
		ReleasePlanID: resp.Entity.ReleasePlanID,
		RepositoryID:  resp.Entity.RepositoryID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(handleAPIError(err))
	}

	return c.JSON(model.APIDataResponse[model.Group]{
		Data: model.Group{
			ID:     resp.Entity.ID,
			Name:   resp.Entity.Name,
			Status: resp.Entity.Status,
		},
	})
}

// @Summary Rmove group API
// @Tags Group
// @Accept json
// @Produce json
// @Param group_id path uuid.UUID true "Group ID"
// @Success 200 {object} model.APIResponse
// @Router /api/v1/groups/{group_id} [delete]
func (h *Group) Remove(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(fmt.Errorf("id is invalid")))
	}
	groupID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(err))
	}

	err = h.groupService.Remove(c.UserContext(), groupID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(err))
	}

	return c.JSON(model.APIResponse{})
}
