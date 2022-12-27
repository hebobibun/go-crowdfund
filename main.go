package main

import (
	"go-crowdfund/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@Bob3214@tcp(localhost:3306)/go_crowdfund?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userSevice := user.NewService(userRepository)
	
	userInput := user.RegisterUserInput{}
	userInput.Name = "Jono"
	userInput.Email = "example@jono.com"
	userInput.Occupation = "Affiliator"
	userInput.Password = "password"

	userSevice.RegisterUser(userInput)
}

// input : input by user
// handler : mapping user input -> struct of input
// service (process/logic) : mapping to struct User 
// repository-> save struct User to DB
// DB