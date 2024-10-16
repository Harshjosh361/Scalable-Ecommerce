package routes

import (
	"harshy/internal/controller"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine, OrderController *controller.OrderController) {

	orderGroup := router.Group("/orders")

	orderGroup.POST("/order", OrderController.CreateOrder)
	orderGroup.GET("/order/:id", OrderController.GetOrderById)
	orderGroup.GET("/order/user/:userid", OrderController.GetOrderByUserId)
}
