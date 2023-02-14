package handlers

import (
	"geek-stash/dtos"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func CreatePlace(db *gorm.DB, context *fiber.Ctx) error {
	place := &dtos.Place{}

	err := context.BodyParser(place)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Request failed",
			"body": nil,
			"status": 422,
		})
		return err
	}

	err = db.Create(place).Error

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Could not create entity",
			"body": nil,
			"status": 500,
		})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Entity Created Successfully",
		"data": nil,
		"status": 200,
	})
	return nil
}


