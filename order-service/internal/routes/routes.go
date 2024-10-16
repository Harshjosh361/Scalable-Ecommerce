package routes

import (
	"harshy/internal/controller"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine, OrderController *controller.OrderController) {

	orderGroup := router.Group("/orders")

	orderGroup.POST("/", OrderController.CreateOrder)
	orderGroup.GET("/:id", OrderController.GetOrderById)
	orderGroup.GET("/user/:userid", OrderController.GetOrderByUserId)
}
