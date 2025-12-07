package main

import (
	"log"

	"github.com/A1fheim/todo-app/internal/config"
	"github.com/A1fheim/todo-app/internal/http/handler"
	"github.com/A1fheim/todo-app/internal/repository"
	"github.com/A1fheim/todo-app/internal/service"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load(".env")

	cfg := config.LoadConfig()

	db, err := repository.NewPostgresPool(
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName)

	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}
	defer db.Close()

	todoRepo := repository.NewTodoPostgres(db)
	todoService := service.NewTodoService(todoRepo)

	h := handler.NewHandler(todoService)
	router := h.InitRoutes()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("error starting server: %s", err)
	}
}
