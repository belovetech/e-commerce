package handlers

import (
	"net/http"

	"github.com/belovetech/e-commerce/database/sqlc"
	"github.com/belovetech/e-commerce/services"
	"github.com/gin-gonic/gin"
)

const (
	ErrInvalidRequest  = "Invalid request data"
	ErrUserExists      = "User already exists"
	ErrHashingPassword = "Error hashing password"
	ErrCreatingUser    = "Error creating user"
	SuccessUserCreated = "User created successfully"
	ErrServer          = "Internal server error"
)

type AuthHandler struct {
	service *services.AuthService
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type RegisterResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func NewAuthHandler(queries *sqlc.Queries) *AuthHandler {
	service := services.NewAuthService(queries)
	return &AuthHandler{service: service}
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrInvalidRequest, "error": err.Error()})
		return
	}

	user, err := h.service.RegisterUser(c, req.Email, req.Password)

	if err != nil {
		switch err {
		case services.ErrUserExists:
			c.JSON(http.StatusConflict, gin.H{"message": ErrUserExists})
		case services.ErrHashingPassword, services.ErrCreatingUser:
			c.JSON(http.StatusInternalServerError, gin.H{"message": ErrServer})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": ErrServer})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": SuccessUserCreated,
		"user": RegisterResponse{
			ID:    int64(user.ID),
			Email: user.Email,
			Role:  user.Role,
		}})
}

func (h *AuthHandler) Login(c *gin.Context) {

	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrInvalidRequest, "error": err.Error()})
		return
	}

	token, err := h.service.LoginUser(c, req.Email, req.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})

}
