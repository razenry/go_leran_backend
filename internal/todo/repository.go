package todo

import "context"

type Repository interface {
	Create(ctx context.Context, todo *Todo) error
	FindAll(ctx context.Context) ([]Todo, error)
	FindByID(ctx context.Context, id uint64) (*Todo, error)
	Update(ctx context.Context, todo *Todo) error
	Delete(ctx context.Context, id uint64) error
}
