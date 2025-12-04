package service

import (
	"context"
	"time"

	"github.com/A1fheim/todo-app/internal/domain/todo"
)

type TodoService interface {
	CreateTodo(ctx context.Context, userID int64, input CreateTodoInput) (todo.Todo, error)
	GetTodoByID(ctx context.Context, userID, id int64) (todo.Todo, error)
	ListTodos(ctx context.Context, userID int64) ([]todo.Todo, error)
	UpdateTodo(ctx context.Context, userID int64, in int64, input UpdateTodoInput) (todo.Todo, error)
	DeleteTodo(ctx context.Context, userID, id int64) error
}

type CreateTodoInput struct {
	Title       string
	Description string
	DueDate     *time.Time
}

type UpdateTodoInput struct {
	Title       *string
	Description *string
	Status      *todo.Status
	DueDate     *time.Time
}
