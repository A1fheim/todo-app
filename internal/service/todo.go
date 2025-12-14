package service

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

type TodoCache interface {
	GetAll(ctx context.Context, userID int64) ([]todo.Todo, error)
	SetAll(ctx context.Context, userID int64, todos []todo.Todo) error
	Invalidate(ctx context.Context, userID int64) error
}

type TodoServiceImpl struct {
	repo  TodoRepository
	cache TodoCache
}

func NewTodoService(repo TodoRepository, cache TodoCache) *TodoServiceImpl {
	return &TodoServiceImpl{repo: repo, cache: cache}
}

func (s *TodoServiceImpl) CreateTodo(ctx context.Context, userID int64, input todo.CreateInput) (todo.Todo, error) {
	t, err := s.repo.Create(ctx, userID, input)
	if err != nil {
		return todo.Todo{}, err
	}
	if s.cache != nil {
		_ = s.cache.Invalidate(ctx, userID)
	}
	return t, nil
}

func (s *TodoServiceImpl) GetTodoByID(ctx context.Context, userID, id int64) (todo.Todo, error) {
	return s.repo.GetByID(ctx, userID, id)
}

func (s *TodoServiceImpl) ListTodos(ctx context.Context, userID int64) ([]todo.Todo, error) {
	if s.cache != nil {
		if todos, err := s.cache.GetAll(ctx, userID); err == nil && todos != nil {
			return todos, nil
		}
	}

	todos, err := s.repo.List(ctx, userID)
	if err != nil {
		return nil, err
	}
	if s.cache != nil {
		_ = s.cache.SetAll(ctx, userID, todos)
	}
	return todos, nil
}

func (s *TodoServiceImpl) UpdateTodo(ctx context.Context, userID, id int64, input todo.UpdateInput) (todo.Todo, error) {
	t, err := s.repo.Update(ctx, userID, id, input)
	if err != nil {
		return todo.Todo{}, err
	}
	if s.cache != nil {
		_ = s.cache.Invalidate(ctx, userID)
	}
	return t, nil
}

func (s *TodoServiceImpl) DeleteTodo(ctx context.Context, userID, id int64) error {
	if err := s.repo.Delete(ctx, userID, id); err != nil {
		return err
	}
	if s.cache != nil {
		_ = s.cache.Invalidate(ctx, userID)
	}
	return nil
}
