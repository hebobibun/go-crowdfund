package main

import (
	"crowdfund/auth"
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
	authService := auth.NewJWT()
	userHandler := handler.NewUserhandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.POST("/email_check", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}
