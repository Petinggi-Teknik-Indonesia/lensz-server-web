package main

import (
	"log"

	"context"
	"lensz-server-web/internal/config"
	"lensz-server-web/internal/router"
	"lensz-server-web/internal/ws"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Run DB seeder
	if err := config.SeedDatabase(context.Background(), db, cfg); err != nil {
		log.Fatalf("❌ Database seeding failed: %v", err)
	}

	// 🔧 Disable automatic redirect for trailing slash
	r := gin.New()
	r.RedirectTrailingSlash = false
	r.Use(gin.Logger(), gin.Recovery())

	// ✅ Global CORS (applies to everything)
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://127.0.0.1:5173",
			"https://www.wandy.web.id",
			"https://wandy.web.id"
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	hub := ws.NewHub()
	go hub.Run()

	router.SetupRoutes(r, db, hub, cfg.JWTSecret)

	// Fallback for OPTIONS (preflight)
	// r.OPTIONS("/*path", func(c *gin.Context) {
	//     c.Status(204)
	// })

	log.Println("🚀 Server running on :" + cfg.ServerPort)
	r.Run("0.0.0.0:" + cfg.ServerPort)
}
