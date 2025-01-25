package controller

import (
	"gcp_go_cloud_run/app/dto"
	"gcp_go_cloud_run/app/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StoreController struct {
	StoreService *usecase.StoreService
}

func NewStoreController(storeService *usecase.StoreService) *StoreController {
	return &StoreController{
		StoreService: storeService,
	}
}

func (sc *StoreController) CreateStore(c *gin.Context) {
	var request struct {
		Name           string `json:"name" binding:"required"`
		DisplayMessage string `json:"displayMessage" binding:"required,max=10"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storeDTO := &dto.StoreDTO{
		Name:           request.Name,
		DisplayMessage: request.DisplayMessage,
	}

	storeID, err := sc.StoreService.CreateStore(storeDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"storeId": storeID})
}

func (sc *StoreController) UpdateDisplayMessage(c *gin.Context) {
	storeID, err := strconv.Atoi(c.Param("storeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid storeId"})
		return
	}

	var request struct {
		DisplayMessage string `json:"displayMessage" binding:"required,max=10"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = sc.StoreService.UpdateDisplayMessage(storeID, request.DisplayMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Display message updated successfully"})
}
