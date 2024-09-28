package controller

import (
	"harsh/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	NotificationService *service.NotificationService
}

func NewNotificationController(notificationservice *service.NotificationService) *NotificationController {
	return &NotificationController{
		NotificationService: notificationservice,
	}
}

func (nc *NotificationController) SendSMS(c *gin.Context) {
	to := c.PostForm("to")
	body := c.PostForm("body")
	err := nc.NotificationService.SendSMS(to, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send notification"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Notification sent successfully"})
}
