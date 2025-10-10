package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/hokkung/release-management-service/config"
	"github.com/hokkung/release-management-service/internal/delivery/rest/model"
	"github.com/hokkung/release-management-service/internal/service/repository"
)

type RepositoryService interface {
	Create(ctx context.Context, req *repository.CreateRequest) error
	Register(ctx context.Context, req *repository.RegisterRequest) error
	Sync(ctx context.Context, req *repository.SyncRequest) error
	List(ctx context.Context, req *repository.ListRequest) (*repository.ListResponse, error)
}

type Repository struct {
	service RepositoryService
	cfg     config.Configuration
}

func NewRepository(service RepositoryService, cfg config.Configuration) *Repository {
	return &Repository{
		service: service,
		cfg:     cfg,
	}
}

// @Summary Register repository API
// @Tags Repository
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Register request"
// @Success 200 {object} RegisterResponse
// @Router /api/v1/repositories/register [post]
func (h *Repository) Register(c *fiber.Ctx) error {
	var req model.RegisterRepositoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(handleAPIError(err))
	}

	err := h.service.Register(c.UserContext(), &repository.RegisterRequest{
		Name:  req.RepositoryNames[0],
		Owner: h.cfg.GitHub.Owner,
	})
	if err != nil {
		return c.JSON(handleAPIError(err))
	}
	return c.JSON(model.APIResponse{})
}

func (h *Repository) List(c *fiber.Ctx) error {
	resp, err := h.service.List(c.UserContext(), &repository.ListRequest{})
	if err != nil {
		return c.JSON(handleAPIError(err))
	}

	repos := make([]model.Repository, 0, len(resp.Entites))
	for _, ent := range resp.Entites {
		repos = append(repos, model.Repository{
			ID:     ent.ID,
			Owner:  ent.Owner,
			Name:   ent.Name,
			Url:    ent.Url,
			Status: ent.Status,
		})
	}
	return c.JSON(&model.APIDataResponse[model.ListRepositoryResponse]{
		Data: model.ListRepositoryResponse{
			Repositories: repos,
		},
	})
}

func (h *Repository) Sync(c *fiber.Ctx) error {
	var req model.SyncRepositoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(handleAPIError(err))
	}
	err := h.service.Sync(c.UserContext(), &repository.SyncRequest{
		RepositoryNames: req.RepositoryNames,
		SyncCommitType:  repository.PullRequestCommitType,
	})
	if err != nil {
		return c.JSON(handleAPIError(err))
	}

	return c.JSON(model.APIResponse{})
}
