package handlers

import (
	"geek-stash/dtos"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateSpecie (db *gorm.DB, context *fiber.Ctx) error {

	specie := &dtos.Specie{}

	err := context.BodyParser(specie)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Unable to process entity",
			"data": nil,
			"status": 422,
		})
		return err
	}

	err = db.Create(specie).Error

	if err !=  nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "An error occured",
			"data": nil,
			"status": 500,
		})
		return err
	}

	context.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "Entity created successfully",
		"data": nil,
		"status": 201,
	})

	return nil
}