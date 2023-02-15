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
			"body":    nil,
			"status":  422,
		})
		return err
	}

	err = db.Create(place).Error

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Could not create entity",
			"body":    nil,
			"status":  500,
		})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Entity Created Successfully",
		"data":    nil,
		"status":  200,
	})
	return nil
}

func GetPlace(db *gorm.DB, context *fiber.Ctx) error {

	page := context.Query("page")
	size := context.Query("size")
	id := context.Query("id")

	result := ResultObject("Place", id)
	query, err := GenerateQueryString(id, page, size, `
			select json_build_object(
				'id', pl.id,
				'name', pl.name,
				'image', pl.image,
				'description', pl.description,
				'franchise', json_build_object(
					'id', fr.id,
					'name', fr.name,
					'start_date', fr.start_date,
					'end_date', fr.end_date,
					'image', fr.image,
					'description', fr.description
				)
			) as result from public.places as pl
			inner join public.franchises as fr on fr.id = pl.franchise
			where pl.id = [id];
	`, `
			select json_agg(json_build_object(
				'id', pl.id,
				'name', pl.name,
				'image', pl.image,
				'description', pl.description,
				'franchise', json_build_object(
					'id', fr.id,
					'name', fr.name,
					'start_date', fr.start_date,
					'end_date', fr.end_date,
					'image', fr.image,
					'description', fr.description
				)
			)) as result from public.places as pl
			inner join public.franchises as fr on fr.id = pl.franchise
			offset [page] limit [size];
	`)
	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Error Occured",
			"data":    nil,
			"status":  500,
		})
		return err
	}

	result_object, err := GetResults(db, result, query)

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": err,
			"data":    nil,
			"status":  500,
		})
		return err
	}

	context.Status(200).JSON(&fiber.Map{
		"message": "Entity Successfully fetched",
		"data":    result_object,
		"status":  200,
	})
	return nil
}
