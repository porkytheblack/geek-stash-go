package handlers

import (
	"geek-stash/dtos"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateSpecie(db *gorm.DB, context *fiber.Ctx) error {

	specie := &dtos.Specie{}

	err := context.BodyParser(specie)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Unable to process entity",
			"data":    nil,
			"status":  422,
		})
		return err
	}

	err = db.Create(specie).Error

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "An error occured",
			"data":    nil,
			"status":  500,
		})
		return err
	}

	context.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "Entity created successfully",
		"data":    nil,
		"status":  201,
	})

	return nil
}

func GetSpecie(db *gorm.DB, context *fiber.Ctx) error {

	page := context.Query("page")
	size := context.Query("size")
	id := context.Query("id")

	result := ResultObject("Place", id)
	query, err := GenerateQueryString(id, page, size, `
		select json_build_object(
			'id', sp.id,
			'name', sp.name,
			'nick_name', sp.nick_name,
			'image', sp.image,
			'description', sp.description,
			'franchise', json_build_object(
				'id', fr.id,
				'name', fr.name,
				'start_date', fr.start_date,
				'end_date', fr.end_date,
				'image', fr.image,
				'description', fr.description
			),
			'place', json_build_object(
				'id', pl.id,
				'name', pl.name,
				'image', pl.image,
				'description', pl.description,

			),
		) from public.species as sp
		inner join public.franchises as fr on fr.id = sp.franchise
		inner join public.places as pl on pl.id = sp.place
		where sp.id = [id];

	`, `
			select json_agg(json_build_object(
				'id', sp.id,
				'name', sp.name,
				'nick_name', sp.nick_name,
				'image', sp.image,
				'description', sp.description,
				'franchise', json_build_object(
					'id', fr.id,
					'name', fr.name,
					'start_date', fr.start_date,
					'end_date', fr.end_date,
					'image', fr.image,
					'description', fr.description
				),
				'place', json_build_object(
					'id', pl.id,
					'name', pl.name,
					'image', pl.image,
					'description', pl.description,
				)
			)) from public.species as sp
			inner join public.franchises as fr on fr.id = sp.franchise
			inner join public.places as pl on pl.id = sp.place
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
