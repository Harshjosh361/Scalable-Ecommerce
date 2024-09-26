package controller

import (
	"harshy/internal/models"
	"harshy/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderController struct {
	OrderService *service.OrderService
}

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{
		OrderService: orderService,
	}
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	// decode the data
	var order *models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	// creating order
	err := oc.OrderService.CreateOrder(c.Request.Context(), order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create order"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (oc *OrderController) GetOrderById(c *gin.Context) {
	id := c.Param("id")

	//convert id to objectid
	ObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to decode id"})
		return
	}

	// call the method
	order, err := oc.OrderService.GetOrderById(c.Request.Context(), ObjectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get order"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (oc *OrderController) GetOrderByUserId(c *gin.Context) {

	userid := c.Param("userid")

	// coverting id to object id
	ObjectId, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	//Get order
	orders, err := oc.OrderService.GetOrderByUserId(c.Request.Context(), ObjectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, orders)
		return
	}
	c.JSON(http.StatusOK, orders)
}
