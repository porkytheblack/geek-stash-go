package main

import (
	"geek-stash/repository"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	repo := repository.InitRepo()

	app := fiber.New()

	repo.SetupRoutes(app)

	app.Listen("0.0.0.0:"+getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		return port
	}
	return port
}

func loadEnv() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Unable to load environmental variables %v", err)
		}
	}
	// if production i.e railway, do nothing
}