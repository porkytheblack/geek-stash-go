package handlers

import (
	"geek-stash/dtos"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func CreateCharacter(db *gorm.DB, context *fiber.Ctx) error {
	character := &dtos.Character{}

	err := context.BodyParser(character)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Unable to process entity",
			"body": nil,
			"status": 400,
		})
		return err
	}

	err = db.Create(character).Error

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "An Error occured while creating the entity",
			"data": nil,
			"status": 500,
		})
		return err
	}

	context.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "Entity created succedully",
		"data": nil,
		"status": 201,
	})

	return nil

}