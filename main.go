package main

import (
	"backend-go-mysql/config"
	"backend-go-mysql/internal/container"
	"backend-go-mysql/routes"
	"backend-go-mysql/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	r := gin.Default()

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET is not set")
	}

	jwtService := utils.NewJwtService(secret)

	app := container.Build(jwtService)

	routes.SetupRoutes(r, app)

	r.Run(":8080")
}
