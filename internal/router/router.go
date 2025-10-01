package router

import (
	"lensz-server-web/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(glassesHandler *handler.GlassesHandler) *gin.Engine {
	r := gin.Default()

	// Glasses routes
	glasses := r.Group("/glasses")
	{
		glasses.POST("/", glassesHandler.Create)
		glasses.GET("/", glassesHandler.GetAll)
		glasses.GET("/:id", glassesHandler.GetByID)
		glasses.PUT("/:id", glassesHandler.Update)
		glasses.DELETE("/:id", glassesHandler.Delete)
	}

	return r
}
