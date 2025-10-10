package model

import (
	"github.com/google/uuid"
)

type RegisterRepositoryRequest struct {
	RepositoryNames []string `json:"repositoryNames"`
}

type RegisterRepositoryResponse struct{}

type ListRepositoryRequest struct {
	*Pagination `json:"pagination"`
}

type ListRepositoryResponse struct {
	Repositories []Repository `json:"repositories"`
}

type Repository struct {
	ID     uuid.UUID `json:"id"`
	Owner  string    `json:"owner"`
	Name   string    `json:"name"`
	Url    string    `json:"url"`
	Status string    `json:"status"`
}

type SyncRepositoryRequest struct {
	RepositoryNames []string `json:"repositoriesNames"`
}

type SyncRepositoryResponse struct {
}
