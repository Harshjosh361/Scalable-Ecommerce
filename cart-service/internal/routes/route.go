package routes

import (
	"harsh/internal/controller"

	"github.com/gin-gonic/gin"
)

func CartRoutes(router *gin.Engine, CartController *controller.CartController) {

	//group the routes
	cartGroup := router.Group("/cart")

	cartGroup.POST("/", CartController.CreateCart)
	cartGroup.PUT("/:userId/items", CartController.AddToCart)
	cartGroup.DELETE("/:userId/items/:productId", CartController.RemoveFromCart)
	cartGroup.GET("/:userId", CartController.GetCart)
	cartGroup.DELETE("/:userId/clear", CartController.ClearCart)
}
