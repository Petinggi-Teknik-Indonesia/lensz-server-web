package main

import (
	"gin/configs"
	"lensz-server-web/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.InitDB()

	r := gin.Default()
	router.SetupRoutes(r, db)

	r.Run(":8080")
}
