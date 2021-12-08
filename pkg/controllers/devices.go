package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/matheusmonte/localDeviceAPI/pkg/models"
)

func FindDevices(c *gin.Context) {
	var devices []models.Device
	models.DB.Find(&devices)
	c.JSON(http.StatusOK, gin.H{"data" : devices})
}

func FindDevice(c *gin.Context) {
	var device models.Device
	if err := models.DB.Where("id = ?", c.Param("id")).First(&device).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": device})
}

func CreateNewDevice(c *gin.Context) {
	var input models.Device
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	models.DB.Create(&input)

	c.JSON(http.StatusOK, gin.H{"data" : input})
}