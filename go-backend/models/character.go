package models

import (
	"fmt"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Character struct {
	ID		uuid.UUID		`gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null" json:"id"`
	Name		*string			`gorm:"type:varchar;default:null" json:"name"`
	Bio		*string			`gorm:"type:varchar;default:null" json:"bio"`
	Attributes	*string		`gorm:"type:varchar;default:null" json:"attributes"`
	Description	*string		`gorm:"type:text;default:null" json:"description"`
	Image		*string		`gorm:"type:text;default:null" json:"image"`
	ExpressiveColor	*string	`gorm:"type:varchar;default:null" json:"expressive_color"`
	CreatedBy	uuid.UUID	`gorm:"type:uuid;" json:"created_by"`
	CreatedOn	time.Time	`gorm:"type:timestamptz;default:now()" json:"created_on"`
	UpdatedOn	time.Time	`gorm:"type:timestamptz;default:now()" json:"updated_on"`
	Status		*string		`gorm:"type:item_status;default:'private'" json:"status"`
	Species		uuid.UUID	`gorm:"type:uuid" json:"species"`
	Weapon		uuid.UUID	`gorm:"type:uuid" json:"weapon"`
	Franchise 	uuid.UUID	`gorm:"type:uuid" json:"franchise"`
	Gadget		[]Gadget	`gorm:"foreignKey:inventor;references:ID" `
	
}

func MigrateCharacter ( db *gorm.DB ) {

	err := db.AutoMigrate(&Character{})


	if err != nil {
		log.Fatalf("Unable to migrate Character, failed with error: %v", err)
	}

	fmt.Println("Successfully Migrated Character!!")

}