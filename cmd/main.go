package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"lensz-server-web/internal/config"
	"lensz-server-web/internal/router"
	"lensz-server-web/internal/ws"
)

func main() {
	// load config
	cfg := config.Load()

	// init DB
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	

	// setup router
	r := gin.Default()
	hub := ws.NewHub()
	go hub.Run()
	router.SetupRoutes(r, db, hub)

	// run server
	r.Run(":" + cfg.ServerPort)
}
