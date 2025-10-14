package router

import (
	"lensz-server-web/internal/handler"
	"lensz-server-web/internal/repository"
	"lensz-server-web/internal/service"
	"lensz-server-web/internal/ws"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB, hub *ws.Hub) {

	// Glasses setup
	glassesRepo := repository.NewGlassesRepository(db)
	glassesService := service.NewGlassesService(glassesRepo)
	glassesHandler := handler.NewGlassesHandler(glassesService)
	scannerHandler := handler.NewScannerHandler(hub)
	dependencyService := service.NewGlassesDependencyService(glassesRepo)
	dependencyHandler := handler.NewGlassesDependencyHandler(dependencyService)

	websocket := r.Group("/scanner")
	{
		websocket.POST("/scan", scannerHandler.Scan)
		websocket.POST("/complete", scannerHandler.CompleteRegistration)
		websocket.POST("/cancel", scannerHandler.CancelRegistration)
	}
	r.GET("/ws", ws.ServeWs(hub))

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

	drawers := api.Group("/drawers")
	{
		drawers.POST("/", dependencyHandler.CreateDrawer)
		drawers.GET("/", dependencyHandler.GetAllDrawers)
		drawers.GET("/:id", dependencyHandler.GetDrawerByID)
	}

	brands := api.Group("/brands")
	{
		brands.POST("/", dependencyHandler.CreateBrand)
		brands.GET("/", dependencyHandler.GetAllBrands)
		brands.GET("/:id", dependencyHandler.GetBrandByID)
	}

	companies := api.Group("/companies")
	{
		companies.POST("/", dependencyHandler.CreateCompany)
		companies.GET("/", dependencyHandler.GetAllCompanies)
		companies.GET("/:id", dependencyHandler.GetCompanyByID)
	}
}
