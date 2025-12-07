package repository

import (
	"context"

	"github.com/A1fheim/todo-app/internal/domain/todo"
)

type TodoRepository interface {
	Create(ctx context.Context, userID int64, input todo.CreateInput) (todo.Todo, error)
	GetByID(ctx context.Context, userID, id int64) (todo.Todo, error)
	List(ctx context.Context, userID int64) ([]todo.Todo, error)
	Update(ctx context.Context, userID, id int64, input todo.UpdateInput) (todo.Todo, error)
	Delete(ctx context.Context, userID, id int64) error
}
