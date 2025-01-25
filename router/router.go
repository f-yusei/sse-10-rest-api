package router

import (
	"gcp_go_cloud_run/app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(storeController *controller.StoreController, bellController *controller.BellController) *gin.Engine {
	router := gin.Default()

	// API v1 group
	api := router.Group("/api/v1")
	{
		// Store routes
		api.POST("/stores", storeController.CreateStore)
		api.PATCH("/stores/:storeId/display_message", storeController.UpdateDisplayMessage)

		// Bell routes
		api.GET("/bells/active", bellController.GetActiveBells)
		//TODO: Uncomment the following routes after implementing the corresponding methods in the BellController
		/*
			api.POST("/bells/:bellId/call", bellController.CallBell)         // Call a bell
			api.POST("/bells/:bellId/complete", bellController.CompleteCall) // Complete a call
		*/
	}

	// Health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	return router
}
