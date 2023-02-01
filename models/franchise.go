package models

import (
	"fmt"
	"log"
	"time"

	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type Franchise struct {
	ID				uuid.UUID		`gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name			string			`gorm:"type:varchar;not null" json:"name"`
	StartDate		*string 		`gorm:"type:varchar;default:null" json:"start_date"`
	EndDate			*string 		`gorm:"type:varchar;default:null" json:"end_date"`	
	Image			*string 		`gorm:"type:text;default:null" json:"image"`
	Description		*string 		`gorm:"type:text;default:null" json:"description"`
	CreatedOn		time.Time		`gorm:"not null;type:timestamptz;default:current_timestamp" json:"created_on"`
	UpdatedOn		time.Time		`gorm:"not null;type:timestamptz;default:current_timestamp" json:"updated_on"`
	CreatedBy		uuid.UUID 		`gorm:"type:uuid; not null" json:"created_by"`
	Status			string			`gorm:"default:'private'" json:"status"`	
	Gadget			[]Gadget		`gorm:"foreignKey:franchise;references:ID"`
	Specie			[]Specie		`gorm:"foreignKey:franchise;references:ID"`
	Character		[]Character		`gorm:"foreignKey:franchise;references:ID"`
}

func MigrateFranchise ( db *gorm.DB ) {
	err := db.AutoMigrate(&Franchise{})

	if err != nil {
		log.Fatalf("Migration Error Occured :: %v", err)
	}

	fmt.Println("Successfull Migrated Franchise!!")

} 