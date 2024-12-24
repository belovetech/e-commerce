package seeder

import (
	"context"
	"log"

	"github.com/belovetech/e-commerce/database/sqlc"
	"github.com/belovetech/e-commerce/utils"
)

func SeedAdminUser(queries *sqlc.Queries) {
	hashedPassword, err := utils.HashPassword("admin@123")
	adminEmail := "admin@example.com"
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	adminExist, err := queries.GetUserByEmail(context.Background(), adminEmail)
	if err != nil {
		log.Fatalf("Failed to get admin user: %v", err)
	}

	if adminExist.Email != "" {
		log.Println("Admin user already exists")
		return
	}

	adminUser := sqlc.CreateUserParams{
		Email:    adminEmail,
		Password: hashedPassword,
		Role:     "admin",
	}

	_, err = queries.CreateUser(context.Background(), adminUser)
	if err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	log.Println("Admin user created successfully")
}
