package repository

import (
	"geek-stash/handlers"
	"geek-stash/middleware"
	"geek-stash/models"
	"geek-stash/storage"
	"log"
	"os"

	// "time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)


type Repository struct {
	DB *gorm.DB
}




func InitRepo () Repository {

	config := &storage.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		SSLMode: os.Getenv("DB_SSLMODE"),
		DBName: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Could not load DB connection")
	}

	// Migrations
	models.RunAllMigrations(db) //migrating franchise


	if err != nil {
		log.Fatal("Unable to migrate")
	}

	return Repository{
		DB: db,
	}
}

func (repo *Repository) SetupRoutes(app *fiber.App){

	app.Use(logger.New())
	app.Use(middleware.Auth)
	api := app.Group("/api")

	GenerateHandlers( handlers.APIHandlers , api, repo.DB)


}


func GenerateHandlers( handlers []handlers.APIHandler, api fiber.Router, db *gorm.DB) {
	
	for _, handler := range handlers {
		_h := func (ctx *fiber.Ctx) error {
			return handler.Handler(db, ctx)
		}
		switch handler.Method {
		case	"GET":
			api.Get(handler.Route, _h)
			log.Printf("Request ::get:: %s done", handler.Route)
		case	"POST":
			api.Post(handler.Route, _h)
			log.Printf("Request ::post:: %s done", handler.Route)
		case	"PUT":
			api.Put(handler.Route, _h)
			log.Printf("Request ::put:: %s done", handler.Route)
		case	"DELETE":
			api.Delete(handler.Route, _h)
			log.Printf("Request ::delete:: %s done", handler.Route)
		default:
			api.Get(handler.Route, _h)
			log.Printf("Request :: %s done", handler.Route)
		}
	}

}