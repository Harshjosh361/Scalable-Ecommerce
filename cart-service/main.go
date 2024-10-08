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

	db := db.Init()

	CartStore := data.NewCartStore(db)
	CartService := service.NewCartService(CartStore)
	CartController := controller.NewCartController(CartService)

	// home route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Cart Server is live",
		})
	})

	// routes
	router.POST("/cart", CartController.CreateCart)
	router.POST("/cart/:userId/items", CartController.AddToCart)
	router.DELETE("/cart/:userId/items/:productId", CartController.RemoveFromCart)
	router.GET("/cart/:userId", CartController.GetCart)
	router.DELETE("/cart/:userId/clear", CartController.ClearCart)

	if err := router.Run(":8085"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
