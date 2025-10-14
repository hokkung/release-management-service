package gorem

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepository[E Entity] struct {
	db             *gorm.DB
	primaryKeyName string
}

func NewBaseRepository[E Entity](db *gorm.DB) *BaseRepository[E] {
	return &BaseRepository[E]{
		db:             db,
		primaryKeyName: getPrimaryKey[E](),
	}
}

func (r *BaseRepository[E]) GetDB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(gormContext).(*gorm.DB)
	if !ok {
		return r.db.WithContext(ctx)
	}
	return tx
}
