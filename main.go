package main

import (
	"crowdfund/config"
	"crowdfund/handler"
	"crowdfund/user"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserhandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.POST("/email_check", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}
