package seeders

import (
	"context"
	"log"
	"strings"

	"github.com/belovetech/e-commerce/config"
	"github.com/belovetech/e-commerce/database/sqlc"
	"github.com/belovetech/e-commerce/utils"
)

type AdminSeeder struct{}

const (
	ErrNoFound = "sql: no rows in result set"
)

func (s AdminSeeder) Name() string {
	return "AdminSeeder"
}

func (s AdminSeeder) Seed(queries *sqlc.Queries, cfg *config.Config) error {
	hashedPassword, err := utils.HashPassword(strings.Trim(cfg.AdminPassword, " "))
	adminEmail := strings.Trim(cfg.AdminEmail, " ")
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	adminExist, err := queries.GetUserByEmail(context.Background(), adminEmail)
	if err.Error() != ErrNoFound {
		log.Fatalf("Something went wrong: %v", err)
	}

	if adminExist.Email != "" {
		log.Println("Admin user already exists")
		return nil
	}

	adminUser := sqlc.CreateUserParams{
		Email:    adminEmail,
		Password: hashedPassword,
		Role:     "admin",
	}

	if _, err = queries.CreateUser(context.Background(), adminUser); err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	log.Println("Admin user created successfully")
	return nil
}
