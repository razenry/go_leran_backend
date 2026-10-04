package todo

import (
	"time"
)

type Todo struct {
	ID          uint64
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Todo) Complete() {
	t.Completed = true
}

func (t *Todo) Reopen() {
	t.Completed = false
}		

func (t *Todo) IsCompleted() bool {
	return t.Completed
}
