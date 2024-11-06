package routes

import (
	"harsh/internal/controller"

	"github.com/gin-gonic/gin"
)

// DefineUserRoutes sets up the routes for user-related operations
func UserRoutes(router *gin.Engine, userController *controller.UserController) {
	// Group user-related routes under the common path "/users"
	userGroup := router.Group("/api/v1/users")

	userGroup.GET("/:id", userController.GetUser)          // Get a user by ID: /users/:id
	userGroup.POST("/register", userController.CreateUser) // Register a new user: /users/register
	userGroup.POST("/login", userController.Login)         // Login a user: /users/login
}
