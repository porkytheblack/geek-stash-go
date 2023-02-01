package handlers

import (
	"geek-stash/dtos"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateProfile (db *gorm.DB, context *fiber.Ctx) error {
	profile := &dtos.Profile{}

	err := context.BodyParser(profile)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Request Faliled",
			"body": nil,
			"status": 422,
		})
		return err
	}

	err = db.Create(profile).Error 

	if err != nil {
		context.Status(http.StatusOK).JSON(&fiber.Map{
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



	return nil;
}