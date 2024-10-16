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

	productStore := data.NewProductStore(db)
	productService := service.NewProductService(productStore)
	productController := controller.NewProductController(productService)

	// Health Check
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Product Server is live",
		})
	})

	// Define product routes using the new routes file
	routes.ProductRoutes(router, productController)

	// starting server
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
