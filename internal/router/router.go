package router

import (
	"lensz-server-web/internal/handler"
	"lensz-server-web/internal/repository"
	"lensz-server-web/internal/service"
	"lensz-server-web/internal/ws"
	"net/http"


	"github.com/gin-contrib/cors"
	"time"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
