package handler

import "github.com/gofiber/fiber/v2"

type GroupHandler struct {}

func NewGroupHandler() *GroupHandler {
	return &GroupHandler{}
}

func (h *GroupHandler) CreateGroup(c *fiber.Ctx) error {
	return nil
}

func (h *GroupHandler) Move(c *fiber.Ctx) error {
	return nil
}

func (h *GroupHandler) UpdateStatus(c *fiber.Ctx) error {
	return nil
}

func (h *GroupHandler) Remove(c *fiber.Ctx) error {
	return nil
}
