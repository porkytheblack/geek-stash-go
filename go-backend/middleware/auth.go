package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func Auth (ctx *fiber.Ctx) error {

	if ctx.Path() != "/api/keys/new" {
		authorization := ctx.GetReqHeaders()["Authorization"]

		if authorization == "" {
			return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
				"message": "Unauthorized",
				"data": nil,
				"status": 401,
			})
			
		} 
		// Bearer 

		tokenString := authorization[7:]
		
		token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("AUTH0_SECRET")), nil
		})
		
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).SendString("Invalid ")
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Printf("Moving to function");
			ctx.Locals("uid", tokenString)
			if err != nil {
				return ctx.Status(http.StatusInternalServerError).SendString("Unauthorized")
			}

			return ctx.Next()
		} else {
			return ctx.Status(401).SendString("Invalid token")
		}


	} else {
		if ctx.GetReqHeaders()["X-Auth-Id"] != "" {
			return ctx.SendStatus(401)
		} else {
			// fix this later
			return ctx.Next()
		}
	}	

}

type nUID struct {
	UID		*uuid.UUID		`json:"uid"`
}


func SetDBSession (db *gorm.DB, ctx *fiber.Ctx) {

	gen_err := db.Connection(func (tx *gorm.DB) error {
		err := tx.Exec(fmt.Sprintf(`select login('%s');`, ctx.Locals("uid"))).Error

		if err != nil {
			log.Printf("An error occured setting session %s", err)
			return err
		}
		r := &nUID{}
		err = tx.Raw(`
		select regexp_replace(nullif(current_setting('request.jwt.claim.sub', true), ''), '"', '', 'g')::uuid as uid;
		`).Scan(r).Error

		if err != nil {
			log.Printf(`An error occured %v`, err)
			return err
		}

		log.Printf(`UID :: %v`, r)
		return nil
	})

	if gen_err != nil {
		log.Printf("An error occured with this transaction:: %s", gen_err)
	}
	

}