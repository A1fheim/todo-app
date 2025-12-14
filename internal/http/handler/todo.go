package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/A1fheim/todo-app/internal/domain/todo"
	"github.com/gin-gonic/gin"
)

type TodoService interface {
	CreateTodo(ctx context.Context, userID int64, input todo.CreateInput) (todo.Todo, error)
	GetTodoByID(ctx context.Context, userID, id int64) (todo.Todo, error)
	ListTodos(ctx context.Context, userID int64) ([]todo.Todo, error)
	UpdateTodo(ctx context.Context, userID, id int64, input todo.UpdateInput) (todo.Todo, error)
	DeleteTodo(ctx context.Context, userID, id int64) error
}

type Handler struct {
	todoService TodoService
}

func NewHandler(todoService TodoService) *Handler {
	return &Handler{
		todoService: todoService,
	}
}

func (h *Handler) createTodo(c *gin.Context) {
	var input todo.CreateInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	userID := int64(1)

	t, err := h.todoService.CreateTodo(c.Request.Context(), userID, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, t)
}

func (h *Handler) listTodos(c *gin.Context) {
	userID := int64(1)

	todos, err := h.todoService.ListTodos(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (h *Handler) getTodoByID(c *gin.Context) {
	userID := int64(1)

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	t, err := h.todoService.GetTodoByID(c.Request.Context(), userID, id)
	if err != nil {
		if errors.Is(err, todo.ErrTodoNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, t)
}

func (h *Handler) updateTodo(c *gin.Context) {
	userID := int64(1)

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input todo.UpdateInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	t, err := h.todoService.UpdateTodo(c.Request.Context(), userID, id, input)
	if err != nil {
		if errors.Is(err, todo.ErrTodoNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, t)
}

func (h *Handler) deleteTodo(c *gin.Context) {
	userID := int64(1)

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = h.todoService.DeleteTodo(c.Request.Context(), userID, id)
	if err != nil {
		if errors.Is(err, todo.ErrTodoNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
