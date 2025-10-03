package domain

import "github.com/google/uuid"

type Repository struct {
	UIDModel

	Name      string
	Url       string
	ServiceID uuid.UUID
}

func (e *Repository) TableName() string {
	return "rms.repositories"
}

func NewRepository() *Repository {
	return &Repository{
		UIDModel: UIDModel{
			ID: uuid.New(),
		},
	}
}
