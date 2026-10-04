package todo

import "context"

type MemoryRepository struct {
	todos  []Todo
	nextID uint64
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		todos:  make([]Todo, 0),
		nextID: 1,
	}
}

var _ Repository = (*MemoryRepository)(nil)

func (r *MemoryRepository) Create(ctx context.Context, item *Todo) error {
	item.ID = r.nextID
	r.nextID++

	r.todos = append(r.todos, *item)

	return nil
}

func (r *MemoryRepository) FindAll(
	ctx context.Context,
) ([]Todo, error) {

	items := make([]Todo, len(r.todos))

	copy(items, r.todos)

	return items, nil
}

// func (r *MemoryRepository) FindAll(
// 	ctx context.Context,
// ) ([]Todo, error) {
// 	return r.todos, nil
// }

func (r *MemoryRepository) FindByID(
	ctx context.Context,
	id uint64,
) (*Todo, error) {

	for i := range r.todos {
		if r.todos[i].ID == id {
			return &r.todos[i], nil
		}
	}

	return nil, ErrTodoNotFound
}

func (r *MemoryRepository) Update(
	ctx context.Context,
	item *Todo,
) error {

	for i := range r.todos {
		if r.todos[i].ID == item.ID {
			r.todos[i] = *item
			return nil
		}
	}

	return ErrTodoNotFound
}

func (r *MemoryRepository) Delete(
	ctx context.Context,
	id uint64,
) error {

	for i := range r.todos {
		if r.todos[i].ID == id {
			r.todos = append(
				r.todos[:i],
				r.todos[i+1:]...,
			)

			return nil
		}
	}

	return ErrTodoNotFound
}
