package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"geek-stash/dtos"
	"geek-stash/models"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ITokenResponse struct {
	AccessToken		string		`json:"access_token"`
	TokenType		string		`json:"token_type"`
}

type IUserProfile struct {
	ID		string		`json:"user_id"`
	UserName	string	`json:"nickname"`
	Email		string	`json:"email"`
	Pic			string	`json:"picture"`
}

func CreateProfile (db *gorm.DB, context *fiber.Ctx) error {
	profile := &dtos.Profile{}
	profileFromDB := &models.Profile{}

	x_auth_key := context.Get("X-Auth-Key")

	if x_auth_key == "" || x_auth_key != os.Getenv("XAUTHKEY") {
		context.Status(http.StatusForbidden).JSON(&fiber.Map{
			"message": "Unauthorized",
			"data": nil,
			"status": 401,
		})
		return errors.New("Unauthorized")
	}


	if err :=  context.BodyParser(profile); err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Unable to process body",
			"body": nil,
			"status": 422,
		})
		return err
	}

	err :=db.First(profileFromDB, "auth_id = ?", profile.Id).Error; 
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			access_token := GetAccessToken(profile)
			if access_token == "" {
				context.Status(http.StatusUnauthorized).JSON(&fiber.Map{
					"message": "Unauthorized",
					"data": nil,
					"status": 401,
				})
				return errors.New("Unauthorized")
			}
			auth0Profile := GetProfileData(access_token, profile.Id)
			if auth0Profile.ID == "" {
				context.Status(http.StatusUnauthorized).JSON(&fiber.Map{
					"message": "Entity does not exist",
					"data": nil,
					"status": 404,
				})
				return errors.New("user doesn't exist")
			}
			log.Printf("Auth Profile:: %v", auth0Profile)
			err:= db.Create(&models.Profile{
				AuthID: auth0Profile.ID,
				UserName: auth0Profile.UserName,
				PicUrl: &auth0Profile.Pic,
			}).Error; 
			if err != nil {
				context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
					"message": "Error creating entity",
					"data": nil,
					"status": 500,
				})
				return err
			}

			context.Status(http.StatusCreated).JSON(fiber.Map{
				"message": "Created Entity Successfully",
				"data": nil,
				"status": 201,
			})
			return nil
		}else {
			context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"message": "Unexpected error",
				"body": nil,
				"status": 500,
			})
			return err
		}
	}

	context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Entity Exists",
		"body": nil,
		"status": 200,
	})

	return nil
}

func GetAccessToken (profileDto *dtos.Profile) (access_token string) {
	url := os.Getenv("AUTH0_ISSUER_BASE_URL")

	audience := url + "/api/v2/"
	url = url + "/oauth/token"

	json_str := fmt.Sprintf(
		"{\"client_id\":\"%s\",\"client_secret\":\"%s\",\"audience\":\"%s\",\"grant_type\":\"%s\"}", 
		os.Getenv("AUTH0_CLIENT_ID"), os.Getenv("AUTH0_CLIENT_SECRET"), audience, "client_credentials")

	log.Println(json_str)
	payload := strings.NewReader(json_str)

	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		log.Printf("An error occured while fetching the Auth:: %s", err)
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("Unable to get access token %s", err )
		return ""
	}

	defer res.Body.Close()

	log.Printf("Status is %v",res.Status)

	content, _ := io.ReadAll(res.Body)

	log.Printf("Body is %v", string(content))
	
	var tokenResponse ITokenResponse;

	if err := json.Unmarshal(content, &tokenResponse); err != nil {
		return ""
	}

	log.Printf("Got access token:: %s", tokenResponse.AccessToken)

	return tokenResponse.AccessToken

}

func GetProfileData (access_token string, id string) *IUserProfile  {
	url := os.Getenv("AUTH0_ISSUER_BASE_URL") + "/api/v2/users/" + id

	log.Printf("Getting user data from: %s", url)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Bearer " + access_token)

	res, err := http.DefaultClient.Do(req)



	if err != nil {
		log.Printf("An error occured %s", err)
		return nil
	}

	log.Printf("Status is:: %v",res.Status)
	
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	log.Printf("Body is %v", string(content))

	var user IUserProfile;

	if err:= json.Unmarshal(content, &user); err != nil {
		return nil
	}

	return &user

}