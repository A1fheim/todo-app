package handler

import (
	"net/http"

	"github.com/A1fheim/todo-app/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	todoService TodoService
	authHandler *AuthHandler
	jwtSecret   string
}

func NewHandler(todoService TodoService, authHandler *AuthHandler, jwtSecret string) *Handler {
	return &Handler{
		todoService: todoService,
		authHandler: authHandler,
		jwtSecret:   jwtSecret,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// auth endpoints (без middleware)
	router.POST("/auth/register", h.authHandler.register)
	router.POST("/auth/login", h.authHandler.login)

	// защищённые маршруты
	todos := router.Group("/todos", middleware.AuthMiddleware(h.jwtSecret))
	{
		todos.POST("/", h.createTodo)
		todos.GET("/", h.listTodos)
		todos.GET("/:id", h.getTodoByID)
		todos.PUT("/:id", h.updateTodo)
		todos.DELETE("/:id", h.deleteTodo)
	}

	return router
}
