package handler

import "github.com/A1fheim/todo-app/internal/service"

type Handler struct {
	todoService service.TodoService
}

func NewHandler(todoService service.TodoService) *Handler {
	return &Handler{
		todoService: todoService,
	}
}
