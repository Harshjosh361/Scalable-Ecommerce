package main

import (
	"harsh/internal/controller"
	"harsh/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	notificationService := service.NewNotificationService()
	notificationController := controller.NewNotificationController(notificationService)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "notification Server is live"})
	})

	router.POST("/message", notificationController.SendSMS)

	router.Run(":8084")
}
