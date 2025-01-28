package main

import (
	"fmt"
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

	envKeys := []string{"DB_HOST", "DB_NAME", "DB_USER", "DB_PASSWORD", "DB_PORT"}
	for _, key := range envKeys {
		value := os.Getenv(key)
		fmt.Printf("%s: %s\n", key, value)
	}

	infrastructure.InitDB()

	r := gin.Default()

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
