package service

import (
	"context"

	"github.com/A1fheim/todo-app/internal/domain/todo"
	"github.com/A1fheim/todo-app/internal/repository"
)

type TodoService interface {
	CreateTodo(ctx context.Context, userID int64, input todo.CreateInput) (todo.Todo, error)
	GetTodoByID(ctx context.Context, userID, id int64) (todo.Todo, error)
	ListTodos(ctx context.Context, userID int64) ([]todo.Todo, error)
	UpdateTodo(ctx context.Context, userID, id int64, input todo.UpdateInput) (todo.Todo, error)
	DeleteTodo(ctx context.Context, userID, id int64) error
}

type TodoServiceImpl struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoServiceImpl {
	return &TodoServiceImpl{repo: repo}
}

func (s *TodoServiceImpl) CreateTodo(ctx context.Context, userID int64, input todo.CreateInput) (todo.Todo, error) {
	return s.repo.Create(ctx, userID, input)
}

func (s *TodoServiceImpl) GetTodoByID(ctx context.Context, userID, id int64) (todo.Todo, error) {
	return s.repo.GetByID(ctx, userID, id)
}

func (s *TodoServiceImpl) ListTodos(ctx context.Context, userID int64) ([]todo.Todo, error) {
	return s.repo.List(ctx, userID)
}

func (s *TodoServiceImpl) UpdateTodo(ctx context.Context, userID, id int64, input todo.UpdateInput) (todo.Todo, error) {
	return s.repo.Update(ctx, userID, id, input)
}

func (s *TodoServiceImpl) DeleteTodo(ctx context.Context, userID, id int64) error {
	return s.repo.Delete(ctx, userID, id)
}
