package main

import (
	"harsh/db"
	"harsh/internal/controller"
	"harsh/internal/data"
	"harsh/internal/routes"
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
	routes.CartRoutes(router, CartController)

	if err := router.Run(":8085"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
