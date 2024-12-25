package services

import (
	"context"
	"errors"

	"github.com/belovetech/e-commerce/database/sqlc"
	"github.com/belovetech/e-commerce/utils"
)

var (
	ErrUserExists      = errors.New("user already exists")
	ErrHashingPassword = errors.New("error hashing password")
	ErrCreatingUser    = errors.New("error creating user")
)

type AuthService struct {
	queries *sqlc.Queries
}

func NewAuthService(queries *sqlc.Queries) *AuthService {
	return &AuthService{queries: queries}
}

func (s *AuthService) RegisterUser(ctx context.Context, email, password string) (*sqlc.CreateUserRow, error) {

	existingUser, err := s.queries.GetUserByEmail(ctx, email)
	if err == nil && existingUser.Email != "" {
		return nil, ErrUserExists
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, ErrHashingPassword
	}

	user, err := s.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:    email,
		Password: hashedPassword,
		Role:     "user",
	})

	if err != nil {
		return nil, ErrCreatingUser
	}

	return &user, nil
}

func (s *AuthService) LoginUser(ctx context.Context, email, password string) (string, error) {
	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = utils.CheckPasswordHash(password, user.Password)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
