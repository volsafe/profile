package routes

import (
	"profile/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Profile routes
	r.POST("/profile/create", handlers.CreateProfile)
	r.GET("/profile/:userID", handlers.GetProfile)
	r.PUT("/profile/update", handlers.UpdateProfile)
	r.GET("/health", handlers.HealthCheck)
	r.DELETE("profile/delete/:userID", handlers.DeleteProfile)
	

	return r
}