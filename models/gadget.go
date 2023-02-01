package models

import (
	"fmt"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type Gadget struct {
	ID			uuid.UUID		`gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name		*string			`gorm:"type:varchar;default:null" json:"name"`
	NickName	*string			`gorm:"type:varchar;default:null" json:"nick_name"`
	Type 		*string			`gorm:"type:gadget_type;default:'weapon'" json:"type"`
	Image		*string			`gorm:"type:text;default:null" json:"image"`
	ExpressiveColor 	*string	`gorm:"type:varchar;default:null" json:"expressive_color"`
	CreatedBy	uuid.UUID		`gorm:"type:uuid" json:"created_by"`
	CreatedOn	time.Time		`gorm:"type:timestamptz;default:now()" json:"created_on"`
	Description	*string			`gorm:"type:text;default:null" json:"description"`
	Status		*string			`gorm:"type:item_status;default:'private'" json:"status"`
	Franchise	uuid.UUID		`gorm:"type:uuid" json:"franchise"`
	UpdatedOn	time.Time		`gorm:"type:timestamptz;default:now()" json:"updated_on"`
	Inventor	uuid.UUID		`gorm:"type:uuid;default:null" json:"inventor"`
}


func MigrateGadgets (db *gorm.DB) {
	err := db.AutoMigrate(&Gadget{})

	if err != nil {
		log.Fatalf("Unable to Migrate Gadgets %v", err)
	}

	fmt.Println("Successfully Migrated Gadgets!!")
}