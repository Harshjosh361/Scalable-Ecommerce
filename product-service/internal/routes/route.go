package routes

import (
	"harsh/internal/controller"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine, productController *controller.ProductController) {
	// Base route for products
	productGroup := router.Group("/api/v1/products")

	// Define individual
	productGroup.GET("/", productController.GetAllProduct)           // Get all products
	productGroup.GET("/:id", productController.GetProduct)           // Get a product by ID
	productGroup.POST("/", productController.CreateProduct)          // Create a new product
	productGroup.DELETE("/:id", productController.DeleteProduct)     // Delete a product by ID
	productGroup.PUT("/update/:id", productController.ModifyProduct) // Update a product by ID
}
