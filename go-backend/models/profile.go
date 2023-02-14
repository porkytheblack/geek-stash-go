package models

import (
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Profile struct {
	ID		uuid.UUID		`gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4()" json:"id"`
	AuthID		string		`gorm:"type:varchar; unique;" json:"auth_id"`
	UserName	string 		`gorm:"type:varchar;not null" json:"username"`
	PicUrl		*string		`gorm:"type:text;default:null" json:"pic_url"`
	Access		string		`gorm:"type:access_level;default:'user'"`
	Franchise 	[]Franchise	`gorm:"foreignKey:CreatedBy;references:ID"`
	Gadget		[]Gadget	`gorm:"foreignKey:CreatedBy;references:ID"`
	Specie		[]Specie	`gorm:"foreignKey:CreatedBy;references:ID"`
	Character	[]Character	`gorm:"foreignKey:CreatedBy;references:ID"`
	Place		[]Place		`gorm:"foreignKey:CreatedBy;references:ID"`
	Keys		[]Keys		`gorm:"foreignKey:Owner;references:ID"`
}

func MigrateProfile (db *gorm.DB) {
	err := db.AutoMigrate(&Profile{})
	if err != nil {
		log.Fatalf("Migration Error Occured :: %v", err)
	}
	fmt.Println("Migration of Profile Successfull!!")
}