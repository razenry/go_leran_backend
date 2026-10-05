package todo

import (
	"context"
	"errors"
	"testing"
)

func TestService_Create(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository)

	ctx := context.Background()

	item, err := service.Create(
		ctx,
		"Belajar Go",
		"Belajar Testing",
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if item.ID != 1 {
		t.Fatalf("expected ID 1, got %d", item.ID)
	}

	if item.Title != "Belajar Go" {
		t.Fatalf(
			"expected title %q, got %q",
			"Belajar Go",
			item.Title,
		)
	}
}

func TestService_FindAll(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository)

	ctx := context.Background()

	item, err := service.Create(
		ctx,
		"Belajar Go",
		"Belajar Testing",
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = service.List(ctx)

	// fmt.Printf("%#v\n", data)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if item.ID != 1 {
		t.Fatalf("expected ID 1, got %d", item.ID)
	}

	if item.Title != "Belajar Go" {
		t.Fatalf(
			"expected title %q, got %q",
			"Belajar Go",
			item.Title,
		)
	}
}

func TestService_Create_InvalidTitle(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository)

	ctx := context.Background()

	_, err := service.Create(
		ctx,
		"",
		"Description",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, ErrInvalidTitle) {
		t.Fatalf(
			"expected ErrInvalidTitle, got %v",
			err,
		)
	}
}

func TestService_Get(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository)

	ctx := context.Background()

	created, err := service.Create(
		ctx,
		"Belajar Go",
		"Repository pattern",
	)

	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	found, err := service.Get(ctx, created.ID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}

	if found.ID != created.ID {
		t.Fatalf(
			"expected ID %d, got %d",
			created.ID,
			found.ID,
		)
	}

	if found.Title != created.Title {
		t.Fatalf(
			"expected title %q, got %q",
			created.Title,
			found.Title,
		)
	}
}

func TestService_Get_NotFound(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository)

	ctx := context.Background()

	_, err := service.Get(ctx, 999)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, ErrTodoNotFound) {
		t.Fatalf(
			"expected ErrTodoNotFound, got %v",
			err,
		)
	}
}

func TestService_Delete(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository)

	ctx := context.Background()

	item, err := service.Create(
		ctx,
		"Todo yang akan dihapus",
		"",
	)

	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	err = service.Delete(ctx, item.ID)

	if err != nil {
		t.Fatalf("delete failed: %v", err)
	}

	_, err = service.Get(ctx, item.ID)

	if !errors.Is(err, ErrTodoNotFound) {
		t.Fatalf(
			"expected ErrTodoNotFound, got %v",
			err,
		)
	}
}

func TestService_Delete_InvalidID(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository)

	err := service.Delete(context.Background(), 0)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, ErrInvalidID) {
		t.Fatalf("expected ErrInvalidID, got %v", err)
	}
}

func TestService_Delete_NotFound(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository)

	err := service.Delete(context.Background(), 999)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, ErrTodoNotFound) {
		t.Fatalf("expected ErrTodoNotFound, got %v", err)
	}
}

func TestService_Get_InvalidID(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository)

	_, err := service.Get(context.Background(), 0)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, ErrInvalidID) {
		t.Fatalf("expected ErrInvalidID, got %v", err)
	}
}

func TestService_List(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository)

	ctx := context.Background()

	_, err := service.Create(ctx, "Todo 1", "")
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	_, err = service.Create(ctx, "Todo 2", "")
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	_, err = service.Create(ctx, "Todo 3", "")
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	items, err := service.List(ctx)

	if err != nil {
		t.Fatalf("list failed: %v", err)
	}

	if len(items) != 3 {
		t.Fatalf("expected 3 todos, got %d", len(items))
	}
}

func TestService_Update(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository)

	ctx := context.Background()

	_, err := service.Create(ctx, "Todo 1", "")
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	err = service.Update(ctx, "Test", "", 1)
	if err != nil {
		t.Fatalf("update failed: %v", err)
	}

	item, err := service.Get(ctx, 1)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}

	if item.Title != "Test" {
		t.Fatalf(
			"expected title %q, got %q",
			"Test",
			item.Title,
		)
	}
}

func TestService_Update_PreservesCompleted(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository)

	ctx := context.Background()

	item, err := service.Create(ctx, "Todo 1", "")
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	err = service.Complete(ctx, item.ID)
	if err != nil {
		t.Fatalf("complete failed: %v", err)
	}

	err = service.Update(ctx, "Todo Updated", "", item.ID)
	if err != nil {
		t.Fatalf("update failed: %v", err)
	}

	item, err = service.Get(ctx, item.ID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}

	if !item.Completed {
		t.Fatal("expected todo to remain completed after update")
	}
}

func TestService_Complete(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository)

	ctx := context.Background()

	item, err := service.Create(ctx, "Todo 1", "")
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	err = service.Complete(ctx, item.ID)
	if err != nil {
		t.Fatalf("complete failed: %v", err)
	}

	item, err = service.Get(ctx, item.ID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}

	if !item.Completed {
		t.Fatal("expected todo to be completed")
	}
}

func TestService_Complete_NotFound(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository)

	ctx := context.Background()

	err := service.Complete(ctx, 999)

	if !errors.Is(err, ErrTodoNotFound) {
		t.Fatalf(
			"expected ErrTodoNotFound, got %v",
			err,
		)
	}
}

func TestService_Complete_InvalidID(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository)

	ctx := context.Background()

	err := service.Complete(ctx, 0)

	if !errors.Is(err, ErrInvalidID) {
		t.Fatalf(
			"expected ErrInvalidID, got %v",
			err,
		)
	}
}
