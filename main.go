package main

import (
	"os"
	"sockets-go/infrastructure"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Iniciar variables de entorno desde .env

	godotenv.Load()

	r := gin.Default()

	infrastructure.SetRoutes(r)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

}
