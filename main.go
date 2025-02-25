package main

import (
	"fmt"
	"gcp_go_cloud_run/app/controller"
	infrastructure "gcp_go_cloud_run/app/infrastructure/mysql"
	mysql "gcp_go_cloud_run/app/infrastructure/mysql/repository"
	"gcp_go_cloud_run/app/router"
	"gcp_go_cloud_run/app/usecase"
	"log"
	"os"
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

	//シード
	infrastructure.Seed()

	// リポジトリの初期化
	storeRepo := mysql.NewStoreRepository(infrastructure.DB)
	bellRepo := mysql.NewBellRepository(infrastructure.DB)

	// サービスの初期化
	storeService := usecase.NewStoreService(storeRepo)
	bellService := usecase.NewBellService(bellRepo)

	// コントローラの初期化
	storeController := controller.NewStoreController(storeService)
	bellController := controller.NewBellController(bellService)

	// ルーターの設定
	r := router.SetupRouter(storeController, bellController)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
