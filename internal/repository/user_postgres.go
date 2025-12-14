package repository

import (
	"context"
	"errors"

	"github.com/A1fheim/todo-app/internal/domain/user"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserPostgres struct {
	db *pgxpool.Pool
}

func NewUserPostgres(db *pgxpool.Pool) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(ctx context.Context, username, passwordHash string) (*user.User, error) {
	query := `
INSERT INTO users (username, password_hash)
VALUES ($1, $2)
RETURNING id, username, password_hash, created_at
`
	var u user.User
	err := r.db.QueryRow(ctx, query, username, passwordHash).Scan(
		&u.ID, &u.Username, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserPostgres) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	query := `
SELECT id, username, password_hash, created_at
FROM users
WHERE username = $1
`
	var u user.User

	err := r.db.QueryRow(ctx, query, username).Scan(
		&u.ID, &u.Username, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrUserNotFound
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserPostgres) GetByID(ctx context.Context, id int64) (*user.User, error) {
	query := `
SELECT id, username, password_hash, created_at
FROM users
WHERE id = $1
`
	var u user.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&u.ID, &u.Username, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrUserNotFound
		}
		return nil, err
	}
	return &u, nil
}
