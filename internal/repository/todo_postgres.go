package repository

import (
	"context"
	"errors"

	"github.com/A1fheim/todo-app/internal/domain/todo"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoPostgres struct {
	db *pgxpool.Pool
}

func NewTodoPostgres(db *pgxpool.Pool) *TodoPostgres {
	return &TodoPostgres{db: db}
}

func (r *TodoPostgres) Create(ctx context.Context, userID int64, input todo.CreateInput) (todo.Todo, error) {
	query := `
		INSERT INTO todos (user_id, title, description, status, due_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id, user_id, title, description, status, due_date, created_at, updated_at
	`
	var t todo.Todo

	err := r.db.QueryRow(ctx, query,
		userID,
		input.Title,
		input.Description,
		"todo",
		input.DueDate,
	).Scan(
		&t.ID,
		&t.UserID,
		&t.Title,
		&t.Description,
		&t.Status,
		&t.DueDate,
		&t.CreatedAt,
		&t.UpdatedAt,
	)

	return t, err
}

func (r *TodoPostgres) GetByID(ctx context.Context, userID, id int64) (todo.Todo, error) {
	query := `
		SELECT id, user_id, title, description, status, due_date, created_at, updated_at
		FROM todos
		WHERE id = $1 AND user_id = $2
	`

	var t todo.Todo

	err := r.db.QueryRow(ctx, query, id, userID).Scan(
		&t.ID, &t.UserID, &t.Title, &t.Description,
		&t.Status, &t.DueDate, &t.CreatedAt, &t.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return todo.Todo{}, todo.ErrTodoNotFound
		}
		return todo.Todo{}, err
	}

	return t, nil
}

func (r *TodoPostgres) List(ctx context.Context, userID int64) ([]todo.Todo, error) {
	query := `
		SELECT id, user_id, title, description, status, due_date, created_at, updated_at
		FROM todos
		WHERE user_id = $1
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]todo.Todo, 0)

	for rows.Next() {
		var t todo.Todo

		err := rows.Scan(
			&t.ID, &t.UserID, &t.Title, &t.Description,
			&t.Status, &t.DueDate, &t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}

func (r *TodoPostgres) Update(ctx context.Context, userID, id int64, input todo.UpdateInput) (todo.Todo, error) {
	query := `
		UPDATE todos
		SET 
			title = COALESCE($1, title),
			description = COALESCE($2, description),
			status = COALESCE($3, status),
			due_date = COALESCE($4, due_date),
			updated_at = NOW()
		WHERE id = $5 AND user_id = $6
		RETURNING id, user_id, title, description, status, due_date, created_at, updated_at
	`

	var t todo.Todo

	err := r.db.QueryRow(
		ctx,
		query,
		input.Title,
		input.Description,
		input.Status,
		input.DueDate,
		id,
		userID,
	).Scan(
		&t.ID, &t.UserID, &t.Title, &t.Description,
		&t.Status, &t.DueDate, &t.CreatedAt, &t.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return todo.Todo{}, todo.ErrTodoNotFound
		}
		return todo.Todo{}, err
	}

	return t, nil
}

func (r *TodoPostgres) Delete(ctx context.Context, userID, id int64) error {
	query := `DELETE FROM todos WHERE id = $1 AND user_id = $2`

	cmd, err := r.db.Exec(ctx, query, id, userID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return todo.ErrTodoNotFound
	}

	return nil
}
