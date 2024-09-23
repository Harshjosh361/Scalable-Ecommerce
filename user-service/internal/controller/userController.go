package controller

import (
	"harsh/internal/models"
	"harsh/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (s *UserController) CreateUser(c *gin.Context) {
	var user models.User
	// Binding json input into struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// calling service layer to create user
	createdUser, err := s.UserService.CreateUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusOK, createdUser)

}

func (s *UserController) GetUser(c *gin.Context) {
	// get user id from Url
	id := c.Param("id")

	// calling service layer to search for user
	user, err := s.UserService.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}
	// sending found user
	c.JSON(http.StatusOK, user)
}

func (s *UserController) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// Bind JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// generate token
	token, err := s.UserService.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unable to generate token"})
		return
	}

	// sending token
	c.JSON(http.StatusOK, gin.H{"token": token})
}
