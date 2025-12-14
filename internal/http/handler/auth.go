package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/A1fheim/todo-app/internal/domain/user"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Register(ctx context.Context, input user.RegisterInput) (*user.User, error)
	Login(ctx context.Context, input user.LoginInput) (string, error)
}

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) register(c *gin.Context) {
	var input user.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	u, err := h.authService.Register(c.Request.Context(), input)
	if err != nil {
		if errors.Is(err, user.ErrUserAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       u.ID,
		"username": u.Username,
	})
}

func (h *AuthHandler) login(c *gin.Context) {
	var input user.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	token, err := h.authService.Login(c.Request.Context(), input)
	if err != nil {
		if errors.Is(err, user.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
