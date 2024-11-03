package main

import (
	"log"
	"backend-3/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r)

	// Start the server
	if err := r.Run(); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
