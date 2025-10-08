package router

import (
	"lensz-server-web/internal/handler"
	"lensz-server-web/internal/repository"
	"lensz-server-web/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Glasses setup
	glassesRepo := repository.NewGlassesRepository(db)
	glassesService := service.NewGlassesService(glassesRepo)
	glassesHandler := handler.NewGlassesHandler(glassesService)

	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	glasses := api.Group("/glasses")
	{
		glasses.POST("/", glassesHandler.Create)
		glasses.GET("/", glassesHandler.GetAll)
		glasses.GET("/:id", glassesHandler.GetByID)
		glasses.PUT("/:id", glassesHandler.Update)
		glasses.DELETE("/:id", glassesHandler.Delete)
	}
}
