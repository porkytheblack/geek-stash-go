package handlers

import (
	"geek-stash/dtos"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateFranchise(db *gorm.DB,context *fiber.Ctx) error {
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

	err = db.Create(franchise).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not create entity", })
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Entity Created Successfully", "data": nil, "status": 200,})

	return nil
}


func GetFranchise  (db *gorm.DB, context *fiber.Ctx) error {

	log.Printf("Get Franchise Called")

	page := context.Query("page")
	size := context.Query("size")
	id	 := context.Query("id")

	result := ResultObject("Franchise", id)
	query, err := GenerateQueryString(id, page, size, `
			select json_build_object(
				'id', id,
				'name', name,
				'start_date', start_date,
				'end_date', end_date,
				'image', image,
				'description', description
			) as result from public.franchises
			where id = [id];
	`, `
				select json_agg(
					json_build_object(
						'id', id,
						'name', name,
						'start_date', start_date,
						'end_date', end_date,
						'image', image,
						'description', description
					)
				) as result from public.franchises
				limit [size]
				offset [page];
	`)
	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Error Occured",
			"data": nil,
			"status": 500,
		})
		return err
	}

	result_object, err := GetResults(db, result, query)


	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": err,
			"data": nil,
			"status": 500,
		})
		return err
	}

	context.Status(200).JSON(&fiber.Map{
		"message": "Entity Successfully fetched",
		"data": result_object,
		"status": 200,
	})
	return nil
}



