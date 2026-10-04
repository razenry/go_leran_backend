package todo

import "errors"

var (
	ErrTodoNotFound = errors.New("todo not found")
	ErrInvalidTitle = errors.New("todo title cannot be empty")
	ErrInvalidID    = errors.New("todo id must be greater than zero")
)
