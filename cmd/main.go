package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"lensz-server-web/internal/config"
	"lensz-server-web/internal/router"
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
	router.SetupRoutes(r, db)

	// run server
	r.Run(":" + cfg.ServerPort)
}
