package handler

import (
	"net/http"

	"github.com/A1fheim/todo-app/internal/service"
	"github.com/gin-gonic/gin"
)

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
