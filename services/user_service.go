package services

import (
	"context"

	"github.com/belovetech/e-commerce/database/sqlc"
)

type UserService struct {
	queries *sqlc.Queries
}

func NewUserService(queries *sqlc.Queries) *UserService {
	return &UserService{queries: queries}
}
func (s *UserService) GetAdmins(ctx context.Context) ([]sqlc.GetAdminsRow, error) {
	admins, err := s.queries.GetAdmins(ctx)
	if err != nil {
		return nil, err
	}

	return admins, nil

}
