package main

import (
	"diary_api/controller"
	"diary_api/database"
	"diary_api/middleware"
	"diary_api/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})	
}

func loadEnv() {
	err := godotenv.Load(".env.local")

	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register",controller.Register)
	publicRoutes.POST("/login",controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry",controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)

	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}