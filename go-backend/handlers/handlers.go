package handlers

import (
	"errors"
	"fmt"
	"geek-stash/dtos"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


type APIHandler struct {
	Route		string
	Handler		func(db *gorm.DB, context *fiber.Ctx) error
	Method		string
}

func Ping (db *gorm.DB,ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{
		"message": "successs",
		"data": nil,
		"status": 200,
	})
}

var APIHandlers []APIHandler = []APIHandler{
	{
		Route: "",
		Handler: Ping,
		Method: "GET",
	},
	{
		Route: "profile",
		Handler: CreateProfile,
		Method: "POST",
	},
	{
		Route: "keys/new",
		Handler: KeyGen,
		Method: "POST",
	},
	{
		Route: "place/create",
		Handler: CreatePlace,
		Method: "POST",
	},
	{
		Route: "character/create",
		Handler: CreateCharacter,
		Method: "POST",
	},
	{
		Route: "specie/create",
		Handler: CreateSpecie,
		Method: "POST",
	},
	{
		Route: "gadget/create",
		Handler: CreateGadgets,
		Method: "POST",
	},
	{
		Route: "gadget/create",
		Handler: CreateFranchise,
		Method: "POST",
	},
}


func GetResults ( db *gorm.DB, result *interface{}, query string ) (interface{}, error ) {
	full_query := fmt.Sprintf(`
		do $$
			%s
		$$;
	`, query)

	err := db.Raw(full_query).Scan(result).Error

	if err != nil {
		log.Printf("An error occured running query:: %s", err)
		return nil, err
	}

	return result, nil
}


func ResultObject ( resultType string, id string ) interface {} {
	
	var r interface{};

	switch resultType {
	case "Franchise":
		if id != "" {
			r = dtos.GetFranchise{}
		} else {
			r = []dtos.GetFranchise{}
		}
	case "Place":
		if id != ""  {
			r = dtos.GetPlace{}
		} else {
			r = []dtos.GetPlace{}
		}
	case "Specie":
		if id != ""  {
			r = dtos.GetSpecie{}
		} else {
			r = []dtos.GetPlace{}
		}
	case "Gadget":
		if id != ""  {
			r = dtos.GetGadget{}
		} else {
			r = []dtos.GetGadget{}
		}
	case "Character":
		if id != ""  {
			r = dtos.GetCharacter{}
		} else {
			r = []dtos.GetGadget{}
		}
	default:
		// do nothing
	}

	return r

}

func GenerateQueryString ( id string, page string, size string, queryOne string, queryMany string ) (*string, error) {
	var str string;
	if id == "" {
		if page == "" || size == "" {
			return nil, errors.New("page or Size is invalid")
		} else {
			str = strings.Replace(queryMany, "[page]", page, 1)
			str = strings.Replace(str, "[size]", size, 1)
		}
	}else {
		str = strings.Replace(queryOne, "[id]", id, 1)
	}
	return &str, nil
}