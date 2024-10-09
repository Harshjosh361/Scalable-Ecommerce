package controller

import (
	"harsh/internal/model"
	"harsh/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartController struct {
	CartService *service.CartService
}

func NewCartController(cartService *service.CartService) *CartController {
	return &CartController{
		CartService: cartService,
	}
}

func (cc *CartController) CreateCart(c *gin.Context) {
	var cart model.Cart

	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	err := cc.CartService.CreateCart(c.Request.Context(), &cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create cart!"})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func (cc *CartController) AddToCart(c *gin.Context) {
	user_id := c.Param("userId")

	var item model.CartItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// coverting id to objectId
	id, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cant convert id to objectId"})
		return
	}

	err = cc.CartService.AddToCart(c.Request.Context(), item, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add to cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "added to cart"})

}

func (cc *CartController) RemoveFromCart(c *gin.Context) {
	userID := c.Param("userId")
	productID := c.Param("productId")

	//convert user id  and productid to objectID
	userObjectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	productObjectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// calling service layer
	err = cc.CartService.RemoveFromCart(c.Request.Context(), userObjectId, productObjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to remove item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "removed from cart"})
}

func (cc *CartController) ClearCart(c *gin.Context) {
	userid := c.Param("userId")

	userObjectId, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	err = cc.CartService.ClearCart(c.Request.Context(), userObjectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to clear cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cart cleared successfully"})
}

func (cc *CartController) GetCart(c *gin.Context) {
	userid := c.Param("userId")

	userObjectId, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	cart, err := cc.CartService.GetCart(c.Request.Context(), userObjectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch cart"})
		return
	}

	c.JSON(http.StatusOK, cart)
}
