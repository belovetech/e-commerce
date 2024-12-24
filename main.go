package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/belovetech/e-commerce/api"
	"github.com/belovetech/e-commerce/config"
	"github.com/belovetech/e-commerce/database/seeder"
	"github.com/belovetech/e-commerce/database/sqlc"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	// initialize database
	db, err := sql.Open("postgres", cfg.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	//  initialize gin router
	router := gin.Default()

	// Setup routes
	api.SetupRoutes(router, db)

	// Seed admin user
	seeder.SeedAdminUser(sqlc.New(db))

	// Run server
	log.Printf("server is running at %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
