package handler

import "github.com/gofiber/fiber/v2"


type ReleasePlanHandler struct {

}

func NewReleasePlanHandler() *ReleasePlanHandler {
	return &ReleasePlanHandler{}
}

func (h *ReleasePlanHandler) List(c *fiber.Ctx) error {
	return nil
}

func (h *ReleasePlanHandler) SetTargetDeployDate(c *fiber.Ctx) error {
	return nil
}
