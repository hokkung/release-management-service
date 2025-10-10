package gorem

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

func (r *BaseRepository[E]) DeleteByID(ctx context.Context, key interface{}) error {
	primaryField := r.primaryKeyName
	resp, err := gorm.G[E](r.GetDB(ctx)).Where(primaryField+" = ?", key).Delete(ctx)
	if err != nil {
		return nil
	}
	if resp != 1 {
		return fmt.Errorf("unable to delete by id: %+v", key)
	}
	return nil
}

func (r *BaseRepository[E]) Delete(ctx context.Context, ent *E) error {
	resp := r.GetDB(ctx).Delete(ent)
	return resp.Error
}
