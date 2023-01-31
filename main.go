package main

import (
	"crowdfund/auth"
	"crowdfund/config"
	"crowdfund/handler"
	"crowdfund/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewJWT()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzUzMTg0ODAsInVzZXJfaWQiOjE1fQ.jQOErI7LyRFA5KGXrNZab9PJlj1zTsEPNy3ykdjo1So")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println(err)
	}

	if token.Valid {
		fmt.Println("VALID")
		fmt.Println("VALID")
		fmt.Println("VALID")
		fmt.Println("VALID")
	} else {
		fmt.Println("INVALID")
		fmt.Println("INVALID")
		fmt.Println("INVALID")
		fmt.Println("INVALID")
	}

	userHandler := handler.NewUserhandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.POST("/email_check", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}
