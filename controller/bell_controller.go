package controller

import (
	"gcp_go_cloud_run/app/dto"
	"gcp_go_cloud_run/app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BellController struct {
	BellService *usecase.BellService
}

func NewBellController(bellService *usecase.BellService) *BellController {
	return &BellController{
		BellService: bellService,
	}
}

func (ctrl *BellController) GetActiveBells(c *gin.Context) {
	activeBellsDTO, err := ctrl.BellService.GetActiveBells()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get active bells"})
		return
	}
	c.JSON(http.StatusOK, activeBellsDTO)
}

func (ctrl *BellController) CreateBell(c *gin.Context) {
	var bellDTO dto.BellDTO
	if err := c.BindJSON(&bellDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	createdBell, err := ctrl.BellService.CreateBell(&bellDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bell"})
		return
	}
	c.JSON(http.StatusCreated, createdBell)
}
