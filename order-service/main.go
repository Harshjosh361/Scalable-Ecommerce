package main

import (
	"harshy/db"
	"harshy/internal/controller"
	"harshy/internal/data"
	"harshy/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := db.Init()

	OrderData := data.NewOrderData(db)
	OrderService := service.NewOrderService(OrderData)
	OrderController := controller.NewOrderController(OrderService)

	// routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Order Server is live",
		})
	})
	router.POST("/order", OrderController.CreateOrder)
	router.GET("/order/:id", OrderController.GetOrderById)
	router.GET("/order/user/:userid", OrderController.GetOrderByUserId)

	// port
	router.Run(":8082")
}
