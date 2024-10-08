package main

import (
	"harsh/db"
	"harsh/internal/controller"
	"harsh/internal/data"
	"harsh/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database := db.Init()

	// all three layers
	PaymentData := data.NewPaymentData(database)
	PaymentService := service.NewPaymentService(PaymentData)
	PaymentController := controller.NewPaymentController(PaymentService)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Message": "Payment server is live"})
	})

	router.POST("/payment", PaymentController.ProcessPayment)
	router.GET("/payment/:id", PaymentController.GetPayment)

	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
