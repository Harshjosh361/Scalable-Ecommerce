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
	// Initialize gin router
	router := gin.Default()

	// initialize db
	db := db.Init()

	// initialize all 3 layers
	userStore := data.NewUserStore(db)
	userService := service.NewUserService(userStore)
	userController := controller.NewUserController(userService)

	// routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "user Server is live",
		})
	})
	router.GET("/users/:id", userController.GetUser)
	router.POST("/register", userController.CreateUser)
	router.POST("/login", userController.Login)

	// starting server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
