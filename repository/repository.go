package repository

import (
	"geek-stash/dtos"
	"geek-stash/models"
	"geek-stash/storage"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/gofiber/fiber/v2"
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
	err = models.MigrateFranchise(db) //migrating franchise

	if err != nil {
		log.Fatal("Unable to migrate")
	}

	return Repository{
		DB: db,
	}
}

func (repo *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	api.Post("franchise/create", repo.CreateFranchise)
	api.Get("franchise/get", repo.GetFranchises)
}

func (repo *Repository) CreateFranchise(context *fiber.Ctx) error {
	franchise := &dtos.Franchise{}
	
	
	err := context.BodyParser(&franchise)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"message": "Request Failed",
				"body": nil,
				"status": "400",
			})
		return err
	}
	franchise.CreatedOn = time.Now().UTC().Format(time.RFC3339)

	err = repo.DB.Create(franchise).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not create entity", })
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Entity Created Successfully", "data": nil, "status": 200,})

	return nil
}

func (repo *Repository) GetFranchises(context *fiber.Ctx) error {
	franchiseModel := &[]models.Franchise{}

	err := repo.DB.Find(franchiseModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Unable to retrieve entities", "data": nil, "status": 400})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Entities recieved successfully", "entities": franchiseModel, "status": 200})

	return nil
}