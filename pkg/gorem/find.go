package gorem

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *BaseRepository[E]) FindAll(ctx context.Context) ([]E, error) {
	var ents []E
	err := r.GetDB(ctx).Find(&ents).Error
	return ents, err
}

func (r *BaseRepository[E]) FindByKey(ctx context.Context, key interface{}) (*E, bool, error) {
	primaryField := r.primaryKeyName
	ent, err := gorm.G[E](r.db).Where(primaryField+" = ?", key).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &ent, true, nil
}

func (r *BaseRepository[E]) FindByName(ctx context.Context, name string) (*E, bool, error) {
	ent, err := gorm.G[E](r.db).Where("name = ?", name).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &ent, true, nil
}

func (r *BaseRepository[E]) FindByFilter(ctx context.Context, filters map[string]any) ([]E, error) {
	query := r.GetDB(ctx).Model(new(E))
	for field, value := range filters {
		query = query.Where(field+" IN ?", value)
	}
	var ents []E
	if err := query.Find(&ents).Error; err != nil {
		return nil, err
	}
	return ents, nil
}
