package main

import (
	"geek-stash/repository"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load environmental variables")
	}
	repo := repository.InitRepo()

	app := fiber.New()

	repo.SetupRoutes(app)

	app.Listen(":8080")
}