package controller

import (
	model "harsh/internal/models"
	"harsh/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentController struct {
	paymentService *service.PaymentService
}

func NewPaymentController(paymentService *service.PaymentService) *PaymentController {
	return &PaymentController{
		paymentService: paymentService,
	}
}

func (pc *PaymentController) ProcessPayment(c *gin.Context) {
	var payment *model.Payment
	err := c.ShouldBindJSON(&payment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	// call the method
	res, err := pc.paymentService.ProcessPayment(c.Request.Context(), payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Payment failed"})
		return
	}

	// send the stripe id if payment completed
	c.JSON(http.StatusAccepted, gin.H{
		"message":   "payment successfull",
		"paymentID": res,
	})
}

func (pc *PaymentController) GetPayment(c *gin.Context) {

	id := c.Param("id")
	// converting to object id
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// invoke the method
	payment, err := pc.paymentService.GetPaymentById(c.Request.Context(), objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch payment"})
		return
	}

	c.JSON(http.StatusOK, payment)
}
