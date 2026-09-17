package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"11val/config"
	"11val/controllers"
	"11val/models"
	"11val/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectDatabase()

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	userController := controllers.UserController{
		DB: db,
	}

	if userController.DB != nil {
		log.Println("User controller created")
	}

	router := gin.Default()

	routes.UserRoutes(router, &userController)

	router.Run(":8082")
}
