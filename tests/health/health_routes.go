package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupHealthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		// Test database connection
		var result int
		if err := db.Raw("SELECT 1").Scan(&result).Error; err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":  "down",
				"message": "Database connection failed",
				"error":   err.Error(),
			})
			return
		}

		// If everything is fine
		c.JSON(http.StatusOK, gin.H{
			"status":  "up",
			"message": "TaskFlow is running smoothly",
			"db":      "connected",
		})
	})
}
