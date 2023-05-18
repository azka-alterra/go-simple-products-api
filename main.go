package main

import (
	"fmt"
	"go-simple-products-api/database"
	"go-simple-products-api/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

const DEFAULT_PORT = "8080"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error when loading env file. Err: %s", err)
	}

	database.InitDB()

	app := echo.New()

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = DEFAULT_PORT
	}

	appPort := fmt.Sprintf(":%s", port)

	app.Logger.Fatal(app.Start(appPort))
}
