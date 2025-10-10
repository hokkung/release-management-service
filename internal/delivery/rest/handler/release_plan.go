package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hokkung/release-management-service/internal/delivery/rest/model"
	"github.com/hokkung/release-management-service/internal/service/release_plan"
)


type ReleasePlan struct {
	service ReleasePlanService
}

func NewReleasePlan(service ReleasePlanService) *ReleasePlan {
	return &ReleasePlan{
		service: service,
	}
}

// @Summary List release plan API
// @Tags ReleasePlan
// @Accept json
// @Produce json
// @Param request body model.ListReleasePlanRequest true "Release plan request"
// @Success 200 {object} model.APIListReleasePlanResponse
// @Router /api/v1/release-plans [post]
func (h *ReleasePlan) List(c *fiber.Ctx) error {
	var req model.ListReleasePlanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(err))
	}

	resp, err := h.service.ListSummary(c.UserContext(), &release_plan.ListSummaryRequest{
		RepositoryIDs: req.RepositoryIDs,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(handleAPIError(err))
	}

	releases := make([]model.ReleasePlan, 0, len(resp.Entities))
	for _, relaseDto := range resp.Entities {
		var groups []model.Group
		for _, groupDto := range relaseDto.Groups {
			var groupItems []model.GroupItem
			for _, groupItemDto := range groupDto.GroupItems {
				groupItems = append(groupItems, model.GroupItem{
					ID:             groupItemDto.ID,
					CommitSHA:      groupItemDto.CommitSHA,
					CommitAuthor:   groupItemDto.CommitAuthor,
					CommitMesssage: groupItemDto.CommitMesssage,
				})
			}
			groups = append(groups, model.Group{
				ID:         groupDto.ID,
				Name:       groupDto.Name,
				Status:     groupDto.Status,
				GroupItems: groupItems,
			})
		}

		var groupItems []model.GroupItem
		for _, groupItemDto := range relaseDto.UnGroupItems {
			groupItems = append(groupItems, model.GroupItem{
				ID:             groupItemDto.ID,
				CommitSHA:      groupItemDto.CommitSHA,
				CommitAuthor:   groupItemDto.CommitAuthor,
				CommitMesssage: groupItemDto.CommitMesssage,
			})
		}

		releases = append(releases, model.ReleasePlan{
			ID:                     relaseDto.ID,
			TargetDeployDate:       relaseDto.TargetDeployDate,
			Note:                   relaseDto.Note,
			LatestTagCommit:        relaseDto.LatestTagCommit,
			LatestMainBranchCommit: relaseDto.LatestMainBranchCommit,
			RepositoryID:           relaseDto.RepositoryID,
			Groups:                 groups,
			UnGroupItems:           groupItems,
		})
	}

	return c.JSON(&model.APIDataResponse[model.ListReleasePlanResponse]{
		Data: model.ListReleasePlanResponse{
			ReleasePlans: releases,
		},
	})
}

// @Summary Update release plan API
// @Tags ReleasePlan
// @Accept json
// @Produce json
// @Param release_plan_id path uuid.UUID true "Release Plan ID"
// @Param request body model.UpdateReleasePlanRequest true "Release plan request"
// @Success 200 {object} model.APIResponse
// @Router /api/v1/release-plans/{release_plan_id}/update [post]
func (h *ReleasePlan) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(fmt.Errorf("id is invalid")))
	}
	releasePlanID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(err))
	}

	var req model.UpdateReleasePlanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(handleAPIError(err))
	}

	if err := h.service.UpdateTargetDeployDateAndNote(c.UserContext(), &release_plan.UpdateTargetDeployDateAndNoteRequest{
		ID:               releasePlanID,
		TargetDeployDate: req.TargetDeployDate,
		Note:             req.Note,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(handleAPIError(err))
	}

	return c.JSON(model.APIResponse{})
}
