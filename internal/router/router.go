package router

import (
	"lensz-server-web/internal/handler"
	"lensz-server-web/internal/repository"
	"lensz-server-web/internal/service"
	"lensz-server-web/internal/ws"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB, hub *ws.Hub) {
	// ✅ Log every incoming request
	r.Use(func(c *gin.Context) {
		log.Printf("HIT: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	// ✅ Test route for sanity check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// ✅ Setup all routes below
	historyRepo := repository.NewHistoryRepository(db)
	historyService := service.NewHistoryService(historyRepo)

	glassesRepo := repository.NewGlassesRepository(db)
	glassesService := service.NewGlassesService(glassesRepo, historyRepo)

	historyHandler := handler.NewHistoryHandler(historyService, glassesService)
	glassesHandler := handler.NewGlassesHandler(glassesService)

	scannerHandler := handler.NewScannerHandler(hub)

	dependencyService := service.NewGlassesDependencyService(glassesRepo)
	dependencyHandler := handler.NewGlassesDependencyHandler(dependencyService)

	// WebSocket
	r.GET("/ws", ws.ServeWs(hub))

	// Scanner
	websocket := r.Group("/scanner")
	{
		websocket.POST("/scan", scannerHandler.Scan)
		websocket.POST("/complete", scannerHandler.CompleteRegistration)
		websocket.POST("/cancel", scannerHandler.CancelRegistration)

		websocket.PATCH("/status", historyHandler.UpdateStatusByRFID)
	}

	// API
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		glasses := api.Group("/glasses")
		{
			glasses.POST("/", glassesHandler.Create)
			glasses.GET("/", glassesHandler.GetAll)
			glasses.GET("/:id", glassesHandler.GetByID)
			glasses.PUT("/:id", glassesHandler.Update)
			glasses.DELETE("/:id", glassesHandler.Delete)

			glasses.GET("/:id/history", historyHandler.GetByGlassesID)
		}

		drawers := api.Group("/drawers")
		{
			drawers.POST("/", dependencyHandler.CreateDrawer)
			drawers.GET("/", dependencyHandler.GetAllDrawers)
			drawers.GET("/:id", dependencyHandler.GetDrawerByID)
			drawers.PUT("/:id", dependencyHandler.UpdateDrawer)
			drawers.DELETE("/:id", dependencyHandler.DeleteDrawer)
		}

		brands := api.Group("/brands")
		{
			brands.POST("/", dependencyHandler.CreateBrand)
			brands.GET("/", dependencyHandler.GetAllBrands)
			brands.GET("/:id", dependencyHandler.GetBrandByID)
			brands.PUT("/:id", dependencyHandler.UpdateBrand)
			brands.DELETE("/:id", dependencyHandler.DeleteBrand)
		}

		companies := api.Group("/companies")
		{
			companies.POST("/", dependencyHandler.CreateCompany)
			companies.GET("/", dependencyHandler.GetAllCompanies)
			companies.GET("/:id", dependencyHandler.GetCompanyByID)
			companies.PUT("/:id", dependencyHandler.UpdateCompany)
			companies.DELETE("/:id", dependencyHandler.DeleteCompany)
		}
	}
}
