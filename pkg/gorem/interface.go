package gorem

import "context"

type Repository[E Entity] interface {
	Create(ctx context.Context, ent *E) error
	Creates(ctx context.Context, ents []*E) error
	FindAll(ctx context.Context) ([]E, error)
	FindByKey(ctx context.Context, key interface{}) (*E, bool, error)
	FindByName(ctx context.Context, name string) (*E, bool, error)
	FindByFilter(ctx context.Context, filters map[string]any) ([]E, error)
	Save(ctx context.Context, ent *E) error
	DeleteByID(ctx context.Context, key interface{}) error
}
