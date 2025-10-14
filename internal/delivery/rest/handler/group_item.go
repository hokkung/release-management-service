package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/delivery/rest/model"
	"github.com/hokkung/release-management-service/internal/service/group"
)



type GroupItem struct {
	groupItemService GroupItemService
	groupService     GroupService
}

func NewGroupItem(groupItemService GroupItemService, groupService GroupService) *GroupItem {
	return &GroupItem{
		groupItemService: groupItemService,
		groupService:     groupService,
	}
}

// @Summary Move group item API
// @Tags GroupItem
// @Accept json
// @Produce json
// @Param group_item_id path uuid.UUID true "Group Item ID"
// @Param request body model.MoveRequest true "Move item request"
// @Success 200 {object} model.APIResponse
// @Router /api/v1/group-items/{group_item_id}/move [post]
func (h *GroupItem) Move(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(errors.New("id must be specified")))
	}
	groupItemUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(err))
	}
	
	var req model.MoveRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(err))
	}

	groups, err := h.groupService.ListByIDs(c.UserContext(), []uuid.UUID{req.ToGroupID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(handleAPIError(err))
	}

	if len(groups) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(errors.New("group entity not found")))
	}

	g := groups[0]
	err = h.groupItemService.Move(c.UserContext(), &group.MoveRequest{
		ToGroupID:   g.ID,
		GroupItemID: groupItemUUID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(handleAPIError(err))
	}

	return c.JSON(model.APIResponse{})
}
