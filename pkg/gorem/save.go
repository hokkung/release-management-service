package gorem

import "context"

func (r *BaseRepository[E]) Save(ctx context.Context, ent *E) error {
	return r.GetDB(ctx).Save(ent).Error
}
