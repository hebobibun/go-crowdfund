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

	userRepository := user.New(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserhandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)

	router.Run()
}
