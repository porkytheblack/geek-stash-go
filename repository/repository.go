package repository

import (
	"geek-stash/dtos"
	"geek-stash/handlers"
	"geek-stash/models"
	"geek-stash/storage"
	"log"
	"net/http"
	"os"
	"strconv"

	// "time"

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
	models.RunAllMigrations(db) //migrating franchise


	if err != nil {
		log.Fatal("Unable to migrate")
	}

	return Repository{
		DB: db,
	}
}

func (repo *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	//franchise
	api.Post("franchise/create", repo.CreateFranchise)
	api.Get("franchise/get", repo.GetFranchises)

	// Profile
	api.Post("profile/create", repo.CreateProfile)
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

	err = repo.DB.Create(franchise).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not create entity", })
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Entity Created Successfully", "data": nil, "status": 200,})

	return nil
}

func (repo *Repository) CreateProfile(context *fiber.Ctx) error {
	return handlers.CreateProfile(repo.DB, context)
}

func (repo *Repository) GetFranchises(context *fiber.Ctx) error {
	franchiseModel := &[]models.Franchise{}
	franchise_id := context.Query("id")
	size, s_err := strconv.Atoi(context.Query("size"))
	page, p_err := strconv.Atoi(context.Query("page"))
	if s_err != nil {
		size = 10
	}
	if p_err != nil {
		page = 0
	}

	var err error;
	if franchise_id == "" {
		err = repo.DB.Limit(size).Offset(page).Find(franchiseModel).Error
	}else{
		err = repo.DB.First(&franchiseModel,"id = ?", franchise_id).Error
	}

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Unable to retrieve entities", "data": nil, "status": 400})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Entities recieved successfully", "entities": franchiseModel, "status": 200})

	return nil
}