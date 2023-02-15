package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"geek-stash/dtos"
	"geek-stash/utils"
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
	{
		Route: "franchise",
		Handler: GetFranchise,
		Method: "GET",
	},
	{
		Route: "characters",
		Handler: GetCharacter,
		Method: "GET",
	},
	{
		Route: "places",
		Handler: GetPlace,
		Method: "GET",
	},
	{
		Route: "species",
		Handler: GetSpecie,
		Method: "GET",
	},
}


func GetResults ( db *gorm.DB, result interface{}, query string ) (interface{}, error ) {

	var d *string;

	err := db.Raw(query).Scan(&d).Error

	if d == nil {
		return nil, errors.New("nothing found")
	}

	if err != nil {
		log.Printf("An error occured running query:: %s", err)
		return nil, err
	}

	log.Printf("Result: %v", d)
	err = json.Unmarshal([]byte(*d), &result)

	if err != nil {
		log.Printf("An error occured unmarshalling query result:: %s", err)
		return nil, err
	}

	return result, nil
}


func ResultObject ( resultType string, id string ) interface {} {

	r := interface{}(nil)

	switch resultType {
	case "Franchise":
		if id != "" {
			r =  dtos.GetFranchise{}
		} else {
			r =  []dtos.GetFranchise{}
		}
	case "Place":
		if id != ""  {
			r = dtos.GetPlace{}
		} else {
			r =  []dtos.GetPlace{}
		}
	case "Specie":
		if id != ""  {
			r = dtos.GetSpecie{}
		} else {
			r = []dtos.GetSpecie{}
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
			r = []dtos.GetCharacter{}
		}
	default:
		// do nothing
	}

	return r

}

func GenerateQueryString ( id string, page string, size string, queryOne string, queryMany string ) (string, error) {
	var str string;
	if id == "" {
		if page == "" || size == "" {
			return "", errors.New("not found")
		} else {
			if !utils.IsValidNumber(page) || !utils.IsValidNumber(size) {
				return "", errors.New("invalid Request")
			}
			str = strings.Replace(queryMany, "[page]", page, 1)
			str = strings.Replace(str, "[size]", size, 1)
		}
	}else {
		if !utils.IsValidUUID(id) {
			return "", errors.New("not found")
		}
		log.Printf("id: %s", id)
		str = strings.Replace(queryOne, "[id]", fmt.Sprintf(`'%s'`, id), 1)
	}
	return str, nil
}