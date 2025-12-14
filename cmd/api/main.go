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

	rdb, err := repository.NewRedisClient(cfg.Redis.Addr)
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	todoRepo := repository.NewTodoPostgres(db)
	todoCache := repository.NewTodoRedis(rdb)
	userRepo := repository.NewUserPostgres(db)

	todoService := service.NewTodoService(todoRepo, todoCache)
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)

	authHandler := handler.NewAuthHandler(authService)
	h := handler.NewHandler(todoService, authHandler, cfg.JWTSecret)

	router := h.InitRoutes()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("error starting server: %s", err)
	}
}
