package main

import (
	"harsh/db"
	"harsh/internal/controller"
	"harsh/internal/data"
	"harsh/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := db.Init()

	productStore := data.NewProductStore(db)
	productService := service.NewProductService(productStore)
	productController := controller.NewProductController(productService)

	// routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Product Server is live",
		})
	})

	router.GET("/products", productController.GetAllProduct)
	router.GET("/product/:id", productController.GetProduct)
	router.POST("/product", productController.CreateProduct)
	router.DELETE("/product/:id", productController.DeleteProduct)
	router.POST("/update-product/:id", productController.ModifyProduct)

	// starting server
	router.Run(":8081")

}
