package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/A1fheim/todo-app/internal/domain/todo"
	"github.com/redis/go-redis/v9"
)

type TodoRedis struct {
	rdb *redis.Client
	ttl time.Duration
}

func NewTodoRedis(rdb *redis.Client) *TodoRedis {
	return &TodoRedis{
		rdb: rdb,
		ttl: 30 * time.Second,
	}
}

func (r *TodoRedis) GetAll(ctx context.Context, userID int64) ([]todo.Todo, error) {
	key := fmt.Sprintf("todos:user:%d", userID)

	val, err := r.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var todos []todo.Todo
	if err := json.Unmarshal([]byte(val), &todos); err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRedis) SetAll(ctx context.Context, userID int64, todos []todo.Todo) error {
	key := fmt.Sprintf("todos:user:%d", userID)

	b, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	return r.rdb.Set(ctx, key, b, 30*time.Second).Err()
}

func (r *TodoRedis) Invalidate(ctx context.Context, userID int64) error {
	key := fmt.Sprintf("todos:user:%d", userID)
	return r.rdb.Del(ctx, key).Err()
}
