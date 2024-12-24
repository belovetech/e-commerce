package main

import (
	"log"

	"github.com/belovetech/e-commerce/api"
	"github.com/belovetech/e-commerce/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	//  initialize gin router
	router := gin.Default()

	// Setup routes
	api.SetupRoutes(router)

	// Run server
	log.Printf("server is running at %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
