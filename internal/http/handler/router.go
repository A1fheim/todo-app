package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	todos := router.Group("/todos")
	{
		todos.POST("/", h.createTodo)
		todos.GET("/", h.listTodos)
		todos.GET("/:id", h.getTodoByID)
	}

	return router
}
