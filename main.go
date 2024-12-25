package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/belovetech/e-commerce/api"
	"github.com/belovetech/e-commerce/config"
	"github.com/belovetech/e-commerce/database/seeders"
	"github.com/belovetech/e-commerce/database/sqlc"
	"github.com/belovetech/e-commerce/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	utils.SetJWTSecretKey(cfg.JWTSecret)

	// initialize database
	db, err := sql.Open("postgres", cfg.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	//  initialize gin router
	router := gin.Default()

	// Setup routes
	api.SetupRoutes(router, db)

	// Run seeders
	if err := seeders.RunSeeders(sqlc.New(db)); err != nil {
		log.Fatalf("Failed to run seeders: %v", err)
	}
	log.Println("All seeders executed successfully")

	// Run server
	log.Printf("server is running at %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
