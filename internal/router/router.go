package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lensz-server-web/internal/handler"
	"lensz-server-web/internal/middleware"
	"lensz-server-web/internal/repository"
	"lensz-server-web/internal/service"
	"lensz-server-web/internal/ws"
	"log"
	"net/http"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB, hub *ws.Hub, jwtSecret string) {
	// ✅ Log every incoming request
	r.Use(func(c *gin.Context) {
		log.Printf("HIT: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	// ✅ Setup all routes below
	historyRepo := repository.NewHistoryRepository(db)
	historyService := service.NewHistoryService(historyRepo)

	glassesRepo := repository.NewGlassesRepository(db)
	glassesService := service.NewGlassesService(glassesRepo, historyRepo)

	historyHandler := handler.NewHistoryHandler(historyService, glassesService, hub)
	glassesHandler := handler.NewGlassesHandler(glassesService)

	scannerHandler := handler.NewScannerHandler(hub)

	dependencyService := service.NewGlassesDependencyService(glassesRepo)
	dependencyHandler := handler.NewGlassesDependencyHandler(dependencyService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, jwtSecret)
	userHandler := handler.NewUserHandler(userService)
	roleRepo := repository.NewRoleRepository(db)
	roleService := service.NewRoleService(roleRepo)
	roleHandler := handler.NewRoleHandler(roleService)

	orgRepo := repository.NewOrganizationRepository(db)
	orgService := service.NewOrganizationService(orgRepo)
	orgHandler := handler.NewOrganizationHandler(orgService)

	scannerRepo := repository.NewScannerRepository(db)
	scannerService := service.NewScannerService(scannerRepo)
	scannerCRUDHandler := handler.NewScannerCRUDHandler(scannerService)

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
		api.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
		// ------------------- USER ROUTES -------------------
		users := api.Group("/users")
		{
			users.POST("/register", userHandler.Register)
			users.POST("/login", userHandler.Login)

			// admin-only routes
			admin := users.Group("/admin")
			admin.Use(middleware.Authenticate(jwtSecret), middleware.RequireRoles(1, 2))
			{
				// admin.POST("/register", userHandler.AdminRegister)
				admin.PATCH("/verify/:email", userHandler.VerifyUser)
				admin.DELETE("/reject/:email", userHandler.CancelUser)
				admin.GET("/unverified", userHandler.GetAllUnverified)
				admin.GET("/verified", userHandler.GetAllVerified)
				admin.GET("/unverified/org/:id", userHandler.GetUnverifiedByOrg)
				admin.GET("/verified/org/:id", userHandler.GetVerifiedByOrg)
			}
		}

		auth := api.Group("/auth")
		{
			auth.GET("/me")
			auth.POST("/refresh", userHandler.Refresh)
		}

		glasses := api.Group("/glasses")
		glasses.Use(middleware.Authenticate(jwtSecret))
		{
			glasses.POST("", glassesHandler.Create)
			glasses.GET("", glassesHandler.GetAll)
			glasses.GET("/:id", glassesHandler.GetByID)
			glasses.PUT("/:id", glassesHandler.Update)
			glasses.DELETE("/:id", middleware.RequireRoles(1, 2), glassesHandler.Delete)

			glasses.GET("/:id/history", historyHandler.GetByGlassesID)
		}

		drawers := api.Group("/drawers")
		drawers.Use(middleware.Authenticate(jwtSecret))
		{
			drawers.POST("", dependencyHandler.CreateDrawer)
			drawers.GET("", dependencyHandler.GetAllDrawers)
			drawers.GET("/:id", dependencyHandler.GetDrawerByID)
			drawers.PUT("/:id", dependencyHandler.UpdateDrawer)
			drawers.DELETE("/:id", middleware.RequireRoles(1, 2), dependencyHandler.DeleteDrawer)
		}

		brands := api.Group("/brands")
		brands.Use(middleware.Authenticate(jwtSecret))
		{
			brands.POST("", dependencyHandler.CreateBrand)
			brands.GET("", dependencyHandler.GetAllBrands)
			brands.GET("/:id", dependencyHandler.GetBrandByID)
			brands.PUT("/:id", dependencyHandler.UpdateBrand)
			brands.DELETE("/:id", middleware.RequireRoles(1, 2), dependencyHandler.DeleteBrand)
		}

		companies := api.Group("/companies")
		companies.Use(middleware.Authenticate(jwtSecret))
		{
			companies.POST("", dependencyHandler.CreateCompany)
			companies.GET("", dependencyHandler.GetAllCompanies)
			companies.GET("/:id", dependencyHandler.GetCompanyByID)
			companies.PUT("/:id", dependencyHandler.UpdateCompany)
			companies.DELETE("/:id", middleware.RequireRoles(1, 2), dependencyHandler.DeleteCompany)
		}
		roles := api.Group("/roles")
		roles.Use(middleware.Authenticate(jwtSecret))
		{
			roles.GET("", roleHandler.GetAll)
			roles.Use(middleware.RequireRoles(1, 2))
			roles.POST("", roleHandler.Create)
			roles.GET("/:id", roleHandler.GetByID)
			roles.PUT("/:id", roleHandler.Update)
			roles.DELETE("/:id", roleHandler.Delete)
		}

		orgs := api.Group("/organizations")
		{
			orgs.GET("", orgHandler.GetAll)
			orgs.Use(middleware.Authenticate(jwtSecret))
			orgs.POST("", orgHandler.Create)
			orgs.GET("/:id", orgHandler.GetByID)
			orgs.PUT("/:id", orgHandler.Update)
			orgs.DELETE("/:id", orgHandler.Delete)
		}
		scanners := api.Group("/scanners")
		scanners.Use(middleware.Authenticate(jwtSecret))
		{
			scanners.POST("", middleware.RequireRoles(1, 2), scannerCRUDHandler.Create)
			scanners.GET("", scannerCRUDHandler.GetAll)
			scanners.GET("/:id", scannerCRUDHandler.GetByID)
			scanners.DELETE("/:id", middleware.RequireRoles(1, 2), scannerCRUDHandler.Delete)
		}

	}
}
