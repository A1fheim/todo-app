package service

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/A1fheim/todo-app/internal/domain/todo"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
)

type InMemoryTodoService struct {
	mu     sync.RWMutex
	nextID int64
	data   map[int64]todo.Todo
}

func NewInMemoryTodoService() *InMemoryTodoService {
	return &InMemoryTodoService{
		nextID: 1,
		data:   make(map[int64]todo.Todo),
	}
}

func (s *InMemoryTodoService) CreateTodo(ctx context.Context, userID int64, input CreateTodoInput) (todo.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()

	id := s.nextID
	s.nextID++

	t := todo.Todo{
		ID:          id,
		UserID:      userID,
		Title:       input.Title,
		Description: input.Description,
		Status:      todo.StatusTodo,
		DueDate:     input.DueDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	s.data[id] = t

	return t, nil
}

func (s *InMemoryTodoService) GetTodoByID(ctx context.Context, userID, id int64) (todo.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	t, ok := s.data[id]
	if !ok || t.UserID != userID {
		return todo.Todo{}, ErrTodoNotFound
	}
	return t, nil
}

func (s *InMemoryTodoService) ListTodos(ctx context.Context, userID int64) ([]todo.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]todo.Todo, 0, len(s.data))
	for _, t := range s.data {
		if t.UserID == userID {
			result = append(result, t)
		}
	}
	return result, nil
}

func (s *InMemoryTodoService) UpdateTodo(ctx context.Context, userID, id int64, input UpdateTodoInput) (todo.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, ok := s.data[id]
	if !ok || t.UserID != userID {
		return todo.Todo{}, ErrTodoNotFound
	}

	if input.Title != nil {
		t.Title = *input.Title
	}

	if input.Description != nil {
		t.Description = *input.Description
	}

	if input.Status != nil {
		t.Status = *input.Status
	}

	if input.DueDate != nil {
		t.DueDate = input.DueDate
	}

	t.UpdatedAt = time.Now()
	s.data[id] = t

	return t, nil
}

func (s *InMemoryTodoService) DeleteTodo(ctx context.Context, userID, id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, ok := s.data[id]
	if !ok || t.UserID != userID {
		return ErrTodoNotFound
	}

	delete(s.data, id)
	return nil
}
