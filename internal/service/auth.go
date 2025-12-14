package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/A1fheim/todo-app/internal/domain/user"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(ctx context.Context, username, passwordHash string) (*user.User, error)
	GetByUsername(ctx context.Context, username string) (*user.User, error)
	GetByID(ctx context.Context, id int64) (*user.User, error)
}
type AuthService struct {
	users     UserRepository
	jwtSecret []byte
}

func NewAuthService(users UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		users:     users,
		jwtSecret: []byte(jwtSecret),
	}
}

type jwtClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func (s *AuthService) Register(ctx context.Context, input user.RegisterInput) (*user.User, error) {
	_, err := s.users.GetByUsername(ctx, input.Username)
	if err == nil {
		return nil, user.ErrUserAlreadyExists
	}
	if !errors.Is(err, user.ErrUserNotFound) {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u, err := s.users.Create(ctx, input.Username, string(hash))
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *AuthService) Login(ctx context.Context, input user.LoginInput) (string, error) {
	u, err := s.users.GetByUsername(ctx, input.Username)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			return "", user.ErrInvalidCredentials
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(input.Password)); err != nil {
		return "", user.ErrInvalidCredentials
	}

	claims := jwtClaims{
		UserID: u.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprint(u.ID),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}
	return signed, nil
}
