package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/A1fheim/todo-app/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	todoService service.TodoService
}

func NewHandler(todoService service.TodoService) *Handler {
	return &Handler{
		todoService: todoService,
	}
}

func (h *Handler) createTodo(c *gin.Context) {
	var input service.CreateTodoInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	userID := int64(1)

	todo, err := h.todoService.CreateTodo(c.Request.Context(), userID, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
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

	todo, err := h.todoService.GetTodoByID(c.Request.Context(), userID, id)
	if err != nil {
		if errors.Is(err, service.ErrTodoNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, todo)
}
