package gorem

import "context"

type BaseRepositoryInt[T any] interface {
	Create(ctx context.Context, ent *T) error
	Creates(ctx context.Context, ents []*T) error
	FindByKey(ctx context.Context, key interface{}) (*T, bool, error)
	FindByName(ctx context.Context, name string) (*T, bool, error)
	FindByFilter(ctx context.Context, filters map[string]any) ([]T, error)
	Save(ctx context.Context, ent *T) error
	DeleteByID(ctx context.Context, key interface{}) error
}
