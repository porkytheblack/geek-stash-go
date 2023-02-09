package models

import (
	"fmt"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Specie struct {
	ID		uuid.UUID		`gorm:"type:uuid;not null;default:uuid_generate_v4();primaryKey" json:"id"`
	Name	*string			`gorm:"type:varchar;default:null" json:"name"`
	NickName	*string		`gorm:"type:varchar;default:null" json:"nick_name"`
	Franchise	uuid.UUID	`gorm:"type:uuid" json:"franchise"`
	CreatedBy	uuid.UUID	`gorm:"type:uuid" json:"created_by"`
	CreatedOn	time.Time	`gorm:"type:timestamptz;default:now()" json:"created_on"`
	UpdatedOn	time.Time	`gorm:"type:timestamptz;default:now()" json:"updated_on"`
	Description	*string		`gorm:"type:text;default:null" json:"description"`
	Image		*string		`gorm:"type:text;default:null" json:"image"`
	Status		*string		`gorm:"type:item_status;default:'private'" json:"status"`
	Place		uuid.UUID	`gorm:"type:uuid" json:"place"`
	Character	[]Character	`gorm:"foreignKey:Species;references:ID"`
}

func MigrateSpecie(db *gorm.DB){

	err := db.AutoMigrate(&Specie{})

	if err != nil {
		log.Fatalf("Unable to Migrate Specie table %v", err)
	}

	fmt.Println("Migrated Species table successfully!!")

}