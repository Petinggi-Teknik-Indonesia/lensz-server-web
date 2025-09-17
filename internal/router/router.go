package router

import (
	"myapp/internal/handler"
	"myapp/internal/repository"
	"myapp/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	api := r.Group("/api")
	{
		api.POST("/register", userHandler.Register)
		api.GET("/hello",fmc.Println("hello"))
	}
}
