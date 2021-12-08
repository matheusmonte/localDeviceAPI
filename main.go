package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusmonte/localDeviceAPI/pkg/models"
	"github.com/matheusmonte/localDeviceAPI/pkg/controllers"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())

	models.ConnectDatabase()
	r.GET("/devices", controllers.FindDevices)
	r.GET("/devices/:id", controllers.FindDevice)
	r.POST("/devices", controllers.CreateNewDevice)

	r.Run(":8080")
}