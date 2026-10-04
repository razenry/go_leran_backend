package todo

import (
	"context"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(
	ctx context.Context,
	title string,
	description string,
) (*Todo, error) {

	// Validasi
	if title == "" {
		return nil, ErrInvalidTitle
	}

	// Buat Todo
	item := &Todo{
		Title:       title,
		Description: description,
	}

	// Create Repository
	err := s.repository.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	// Return todo
	return item, nil
}

func (s *Service) List(
	ctx context.Context,
) ([]Todo, error) {
	// Return data

	return s.repository.FindAll(ctx)
}

func (s *Service) Get(
	ctx context.Context,
	id uint64,
) (*Todo, error) {

	// Validasi ID
	if id == 0 {
		return nil, ErrInvalidID
	}

	return s.repository.FindByID(ctx, id)
}

func (s *Service) Delete(
	ctx context.Context,
	id uint64,
) error {

	// Validasi ID
	if id == 0 {
		return ErrInvalidID
	}

	return s.repository.Delete(ctx, id)
}

func (s *Service) Update(
    ctx context.Context,
    newTitle string,
    newDesc string,
    id uint64,
) error {
    if newTitle == "" {
        return ErrInvalidTitle
    }

    if id == 0 {
        return ErrInvalidID
    }

    item, err := s.repository.FindByID(ctx, id)
    if err != nil {
        return err
    }

    item.Title = newTitle
    item.Description = newDesc

    return s.repository.Update(ctx, item)
}

func (s *Service) Complete(
	ctx context.Context,
	id uint64,
) error {
	if id == 0 {
		return ErrInvalidID
	}

	item, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	item.Complete()

	return s.repository.Update(ctx, item)
}

// func (s *Service) List(ctx context.Context) ([]Todo, error) {

// 	// FindAll Data
// 	// data, err := s.repository.FindAll(ctx)

// 	// Error Handling
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// return data, nil

// 	// Simple Return
// 	return s.repository.FindAll(ctx)
// }

// func (s *Service) Get(ctx context.Context, id uint64) (*Todo, error) {

// 	// validasi
// 	if id == 0 {
// 		return nil, ErrInvalidID
// 	}

// 	// Get Data By ID
// 	data, err := s.repository.FindByID(ctx, id)

// 	if err != nil {
// 		return nil, ErrTodoNotFound
// 	}

// 	return data, nil
// }

// func (s *Service) Get(
// 	ctx context.Context,
// 	id uint64,
// ) (*Todo, error) {

// 	if id == 0 {
// 		return nil, ErrInvalidID
// 	}

// 	data, err := s.repository.FindByID(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

// func (s *Service) Delete(
// 	ctx context.Context,
// 	id uint64,
// ) error {

// 	// Validasi
// 	if id == 0 {
// 		return ErrInvalidID
// 	}

// 	// Get Data By ID
// 	err := s.repository.Delete(ctx, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
