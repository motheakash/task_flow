package routes

import (
	"task_flow/internal/apps/users/controllers"
	"task_flow/internal/apps/users/repositories"
	"task_flow/internal/apps/users/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userController.CreateUser) // Create User API
	}
}
