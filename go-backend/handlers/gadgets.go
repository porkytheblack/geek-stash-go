package handlers

import (
	"geek-stash/dtos"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateGadgets ( db *gorm.DB, context *fiber.Ctx) error {
	gadget := &dtos.Gadgets{}

	err := context.BodyParser(gadget) 

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Unable to parse entity",
			"data": nil,
			"status": 400,
		})
		return err
	}

	err = db.Create(gadget).Error

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Internal Server Error",
			"data": nil,
			"status":  500,
		})

		return err
	}

	context.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "Entity Created successfully",
		"data": nil,
		"status": 201,
	})

	return nil

}