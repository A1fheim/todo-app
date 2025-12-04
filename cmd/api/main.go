package main

import (
	"log"

	"github.com/A1fheim/todo-app/internal/http/handler"
	"github.com/A1fheim/todo-app/internal/service"
)

func main() {
	todoService := service.NewInMemoryTodoService()

	h := handler.NewHandler(todoService)
	router := h.InitRoutes()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("error starting server: %s", err)
	}
}
