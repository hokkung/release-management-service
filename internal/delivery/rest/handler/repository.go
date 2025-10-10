package handler

import "github.com/gofiber/fiber/v2"

type Repository struct {
}

type RegisterRequest struct{}
type RegisterResponse struct{}

// @Summary Register repository API
// @Tags Repository
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Register request"
// @Success 200 {object} RegisterResponse
// @Router /api/v1/repositories/register [post]
func (h *Repository) Register(c *fiber.Ctx) error {
	return c.JSON(&RegisterResponse{})
}
