package repository

import (
	"context"
	"encoding/json"
	"errors"
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

func (r *TodoRedis) key(userID int64) string {
	return fmt.Sprintf("todos:user:%d", userID)
}

func (r *TodoRedis) GetAll(ctx context.Context, userID int64) ([]todo.Todo, error) {

	val, err := r.rdb.Get(ctx, r.key(userID)).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var todos []todo.Todo
	if err := json.Unmarshal([]byte(val), &todos); err != nil {
		_ = r.rdb.Del(ctx, r.key(userID)).Err()
		return nil, nil
	}
	return todos, nil
}

func (r *TodoRedis) SetAll(ctx context.Context, userID int64, todos []todo.Todo) error {

	b, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	return r.rdb.Set(ctx, r.key(userID), b, r.ttl).Err()
}

func (r *TodoRedis) Invalidate(ctx context.Context, userID int64) error {
	return r.rdb.Del(ctx, r.key(userID)).Err()
}
