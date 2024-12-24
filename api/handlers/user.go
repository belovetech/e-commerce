package handlers

import (
	"github.com/belovetech/e-commerce/database/sqlc"
	"github.com/belovetech/e-commerce/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(queries *sqlc.Queries) *UserHandler {
	service := services.NewUserService(queries)
	return &UserHandler{service: service}
}
