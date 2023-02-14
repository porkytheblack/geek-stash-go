package handlers

import (
	"errors"
	"geek-stash/models"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)


func generateJwt ( uuid *string ) (*string, error) {

	if uuid == nil {
		return nil, errors.New("empty uuid")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = uuid;
	claims["authorized"] = true;
	key := []byte(os.Getenv("AUTH0_SECRET"))
	tokenString, err := token.SignedString(key)

	if err != nil {
		log.Printf("Unable to sign jwt:: %s ", err)
		return nil, err
	}

	return &tokenString, nil
}

func KeyGen ( db *gorm.DB, ctx *fiber.Ctx ) error {

	log.Printf("%v", ctx.GetReqHeaders())

	auth_id := ctx.GetReqHeaders()["X-Auth-Id"]

	if auth_id == "" {
		ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Unauthorized",
			"data": nil,
			"status": 200,
		})
		return errors.New("auth ID is empty, request unauthorized")
	}

	profile := &models.Profile{}

	err := db.Model(&models.Profile{}).First(profile, "auth_id = ?", auth_id).Error

	if err != nil {
		ctx.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "User not found",
			"data": nil,
			"status": 404,
		})
		return err
	}
	str := profile.ID.String()
	token, terr := generateJwt(&str)

	if terr != nil {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": terr,
			"data": nil,
			"status": 500,
		})
		return terr
	}

	key := &models.Keys{
		Key: *token,
		Owner: profile.ID,
	}

	err = db.Create(key).Error

	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": err,
			"data": nil,
			"status": 500,
		})
		return err
	}

	ctx.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "Success",
		"data": nil,
		"status": 201,
	})

	

	return nil

}