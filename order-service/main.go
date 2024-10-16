package main

import (
	"harshy/db"
	"harshy/internal/controller"
	"harshy/internal/data"
	"harshy/internal/routes"
	"harshy/internal/service"
	"log"
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

	routes.OrderRoutes(router, OrderController)

	// port
	if err := router.Run(":8083"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
