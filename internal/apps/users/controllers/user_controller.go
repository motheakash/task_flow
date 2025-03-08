package controllers

import (
	"log"
	"net/http"
	"task_flow/internal/apps/users/models"
	"task_flow/internal/apps/users/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.UserModel

	// Bind request body to user model
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	if err := ctrl.Service.CreateUser(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}
