package main

import (
	"log"

	"github.com/A1fheim/todo-app/internal/http/handler"
)

func main() {
	h := handler.NewHandler()
	router := h.InitRoutes()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("error starting server: %s", err)
	}
}
