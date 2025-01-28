package main

import (
	infrastructure "gcp_go_cloud_run/app/infrastructure/mysql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	infrastructure.InitDB()

	r := gin.Default()

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
