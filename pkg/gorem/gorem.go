package gorem

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type gormKey string

const gormContext gormKey = "gormContext"

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{
		db: db,
	}
}

func (r *BaseRepository[T]) Create(ctx context.Context, ent *T) error {
	err := gorm.G[T](r.db).Create(ctx, ent)
	if err != nil {
		return nil
	}
	return nil
}

func (r *BaseRepository[T]) Creates(ctx context.Context, ents []*T) error {
	err := r.GetDB(ctx).Create(ents)
	if err != nil {
		return nil
	}
	return nil
}

func (r *BaseRepository[T]) GetDB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(gormContext).(*gorm.DB)
	if !ok {
		return r.db.WithContext(ctx)
	}
	return tx
}

func (r *BaseRepository[T]) FindByKey(ctx context.Context, key interface{}) (*T, bool, error) {
	ent, err := gorm.G[T](r.db).Where("id = ?", key).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &ent, true, nil
}

func (r *BaseRepository[T]) FindByName(ctx context.Context, name string) (*T, bool, error) {
	ent, err := gorm.G[T](r.db).Where("name = ?", name).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &ent, true, nil
}

func (r *BaseRepository[T]) Save(ctx context.Context, ent *T) error {
	return r.GetDB(ctx).Save(ent).Error
}

func (r *BaseRepository[T]) DeleteByID(ctx context.Context, key interface{}) error {
	resp, err := gorm.G[T](r.GetDB(ctx)).Where("id = ?", key).Delete(ctx)
	if err != nil {
		return nil
	}
	if resp != 1 {
		return fmt.Errorf("unable to delete by id: %+v", key)
	}
	return nil
}

func (r *BaseRepository[T]) Delete(ctx context.Context, ent *T) error {
	resp := r.GetDB(ctx).Delete(ent)
	return resp.Error
}

func (r *BaseRepository[T]) FindByFilter(ctx context.Context, filters map[string]any) ([]T, error) {
	query := r.GetDB(ctx).Model(new(T))
	var ents []T
	for field, value := range filters {
		query = query.Where(field+" IN ?", value)
	}
	if err := query.Find(&ents).Error; err != nil {
		return nil, err
	}
	return ents, nil
}
