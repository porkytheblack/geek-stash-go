package models

import (
	"fmt"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Place struct {
	ID		uuid.UUID		`gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name	*string			`gorm:"type:varchar;default:null" json:"name"`
	Franchise	uuid.UUID	`gorm:"type:uuid" json:"franchise"`
	CreatedBy	uuid.UUID	`gorm:"type:uuid" json:"created_by"`
	CreatedOn	time.Time	`gorm:"type:timestamptz;default:now()" json:"created_on"`
	UpdatedOn	time.Time	`gorm:"type:timestamptz;default:now()" json:"updated_on"`
	Description	*string		`gorm:"type:text;default:null" json:"description"`
	Image		*string		`gorm:"type:text;default:null" json:"image"`
	Status		*string		`gorm:"type:item_status;default:'private'" json:"status"`
	Specie		[]Specie	`gorm:"foreignKey:place;references:ID"`
}


func MigratePlace ( db *gorm.DB) {

	err := db.AutoMigrate(&Place{})

	if err != nil {
		log.Fatalf("Unable to Migrate Place because of error:: %v", err)
	}

	fmt.Println("Migrated Place Successfully!!")

}