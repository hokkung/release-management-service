package repository

import "github.com/hokkung/release-management-service/internal/domain"

type CreateRequest struct {
	Name  string
	Url   string
	Owner string
}

type RegisterRequest struct {
	Owner string
	Name  string
}

type SyncCommitType string

const NoParentSyncCommitType SyncCommitType = "NoParent"

type SyncRequest struct {
	RepositoryName string
	SyncCommitType SyncCommitType
}

type ListRequest struct {
}

type ListResponse struct {
	Entites []domain.Repository
}
