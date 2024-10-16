package api

import (
	"github.com/gin-gonic/gin"
	"github.com/timothygan/cafewhere/backend/internal/api/handlers"
	"github.com/timothygan/cafewhere/backend/internal/api/middleware"
	"go.uber.org/zap"
)

func SetupRoutes(r *gin.Engine, h *handlers.CafeHandler, logger *zap.Logger) {
	// Add logging middleware
	r.Use(middleware.Logging(logger))

	// Health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	// API routes
	api := r.Group("/api")
	{
		// Coffee shop routes
		cafes := api.Group("/cafes")
		{
			cafes.GET("/search", h.SearchCafes)
			cafes.GET("/details", h.GetCafeDetails)
		}

		// You can add more route groups here for other features
		// For example:
		// users := api.Group("/users")
		// {
		//     users.POST("/register", h.RegisterUser)
		//     users.POST("/login", h.LoginUser)
		// }
	}
}
