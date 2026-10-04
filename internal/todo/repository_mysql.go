package todo

import (
	"context"
	"database/sql"
	"errors"
)

type MySQLRepository struct {
	db *sql.DB
}

var _ Repository = (*MySQLRepository)(nil)

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (r *MySQLRepository) Create(
	ctx context.Context,
	item *Todo,
) error {
	query := `
		INSERT INTO todos (
			title,
			description,
			completed
		)
		VALUES (?, ?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		item.Title,
		item.Description,
		item.Completed,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	item.ID = uint64(id)

	return nil
}

func (r *MySQLRepository) FindAll(
	ctx context.Context,
) ([]Todo, error) {

	query := `
		SELECT
			id,
			title,
			description,
			completed,
			created_at,
			updated_at
		FROM todos
		ORDER BY id ASC 
	`

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := make([]Todo, 0)

	for rows.Next() {
		var item Todo

		err := rows.Scan(
			&item.ID,
			&item.Title,
			&item.Description,
			&item.Completed,
			&item.CreatedAt,
			&item.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, item)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil

}

func (r *MySQLRepository) FindByID(
	ctx context.Context,
	id uint64,
) (*Todo, error) {
	query := `
		SELECT 
			id,
			title,
			description,
			completed,
			created_at,
			updated_at
		FROM todos 
		WHERE id = ?
	`

	var item Todo

	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&item.ID,
		&item.Title,
		&item.Description,
		&item.Completed,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrTodoNotFound
		}

		return nil, err
	}

	return &item, nil

}

func (r *MySQLRepository) Update(
	ctx context.Context,
	item *Todo,
) error {
	query := `
		UPDATE todos
		SET 
			title = ?,
			description = ?,
			completed = ?
		WHERE id = ?
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		item.Title,
		item.Description,
		item.Completed,
		item.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrTodoNotFound
	}

	return nil

}

func (r *MySQLRepository) Delete(
	ctx context.Context,
	id uint64,
) error {

	query := `
		DELETE FROM todos
		WHERE id = ?
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrTodoNotFound
	}

	return nil
}
