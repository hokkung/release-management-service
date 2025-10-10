package gorem

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Entity interface {
	TableName() string
	PrimaryKey() string
}

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type UIDModel struct {
	ID uuid.UUID `gorm:"primaryKey"`
	Model
}
