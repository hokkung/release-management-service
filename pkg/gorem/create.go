package gorem

import (
	"context"

	"gorm.io/gorm"
)

func (r *BaseRepository[E]) Create(ctx context.Context, ent *E) error {
	err := gorm.G[E](r.db).Create(ctx, ent)
	if err != nil {
		return nil
	}
	return nil
}

func (r *BaseRepository[E]) Creates(ctx context.Context, ents []*E) error {
	err := r.GetDB(ctx).Create(ents)
	if err != nil {
		return nil
	}
	return nil
}
