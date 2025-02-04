package controller

import (
	"gcp_go_cloud_run/app/dto"
	"gcp_go_cloud_run/app/usecase"
	"net/http"
	"strconv"

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

func (ctrl *BellController) CallBell(c *gin.Context) {
	bellID := c.Param("bellId")
	if bellID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bell ID"})
		return
	}

	bellIDtoInt, err := strconv.Atoi(bellID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bell ID"})
		return
	}

	err = ctrl.BellService.CallBell(bellIDtoInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call bell"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bell called successfully"})
}

func (ctrl *BellController) CompleteCall(c *gin.Context) {
	bellID := c.Param("bellId")
	if bellID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bell ID"})
		return
	}

	bellIDtoInt, err := strconv.Atoi(bellID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bell ID"})
		return
	}

	err = ctrl.BellService.CompleteCall(bellIDtoInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to complete call"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Call completed successfully"})
}
