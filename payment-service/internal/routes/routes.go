package routes

import (
	"harsh/internal/controller"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(router *gin.Engine, PaymentController *controller.PaymentController) {

	paymentGroup := router.Group("/payment")

	paymentGroup.POST("/", PaymentController.ProcessPayment)
	paymentGroup.GET("/:id", PaymentController.GetPayment)
}
