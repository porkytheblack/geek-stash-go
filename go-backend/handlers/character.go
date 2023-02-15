package handlers

import (
	"geek-stash/dtos"
	"log"
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
			"body":    nil,
			"status":  400,
		})
		return err
	}

	err = db.Create(character).Error

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "An Error occured while creating the entity",
			"data":    nil,
			"status":  500,
		})
		return err
	}

	context.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "Entity created succedully",
		"data":    nil,
		"status":  201,
	})

	return nil

}

func GetCharacter(db *gorm.DB, context *fiber.Ctx) error {
	log.Printf("GetCharacter() called")
	page := context.Query("page")
	size := context.Query("size")
	id := context.Query("id")

	result := ResultObject("Character", id)
	query, err := GenerateQueryString(id, page, size, `
			select json_build_object(
				'id', public.characters.id,
				'name', public.characters.name,
				'image', public.characters.image,
				'description', public.characters.description,
				'bio', bio,
				'attributes', attributes,
				'status', public.characters.status,
				'franchise', json_build_object(
					'id', public.franchises.id,
					'name', public.franchises.name,
					'start_date', public.franchises.start_date,
					'end_date', public.franchises.end_date,
					'image', public.franchises.image,
					'description', public.franchises.description
				),
				'species', json_build_object(
					'id', public.species.id,
					'name', public.species.name,
					'image', public.species.image,
					'description', public.species.description
				),
				'weapon', json_build_object(
					'id', public.gadgets.id,
					'name', public.gadgets.name,
					'image', public.gadgets.image,
					'description', public.gadgets.description
				)
			) as result from public.characters
			inner join public.franchises on public.characters.franchise = public.franchises.id
			inner join public.species on public.characters.species = public.species.id
			inner join public.gadgets on public.characters.weapon = public.gadgets.id
			where public.characters.id = [id];
	`, `
				select json_agg(
					json_build_object(
						'id', public.characters.id,
						'name', public.characters.name,
						'image', public.characters.image,
						'description', public.characters.description,
						'bio', bio,
						'attributes', attributes,
						'status', public.characters.status,
						'franchise', json_build_object(
							'id', public.franchises.id,
							'name', public.franchises.name,
							'start_date', public.franchises.start_date,
							'end_date', public.franchises.end_date,
							'image', public.franchises.image,
							'description', public.franchises.description
						),
						'species', json_build_object(
							'id', public.species.id,
							'name', public.species.name,
							'image', public.species.image,
							'description', public.species.description
						),
						'weapon', json_build_object(
							'id', public.gadgets.id,
							'name', public.gadgets.name,
							'image', public.gadgets.image,
							'description', public.gadgets.description
						)
					)
				) as result from public.characters
				inner join public.franchises on public.characters.franchise = public.franchises.id
				inner join public.species on public.characters.species = public.species.id
				inner join public.gadgets on public.characters.weapon = public.gadgets.id
				limit [size]
				offset [page];
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
