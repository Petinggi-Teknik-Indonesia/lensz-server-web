package main

import (
	"lensz-server-web/config"
	"lensz-server-web/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load env first
	configs.LoadEnv()

	// Setup DB
	db := configs.InitDB()

	// Setup Gin
	r := gin.Default()
	router.SetupRoutes(r, db)

	// Run app on port from env
	port := configs.GetEnv("APP_PORT", "8080")
	r.Run(":" + port)
}
