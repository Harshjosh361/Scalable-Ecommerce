package main

import (
	"harsh/internal/controller"
	"harsh/internal/service"
	"log"
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

	router.POST("/api/v1/notification", notificationController.SendSMS)

	if err := router.Run(":8084"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
