package models

import (
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type Usage struct {
	ID		uuid.UUID		`gorm:"type:uuid;unique;not null;primaryKey;default:uuid_generate_v4();" json:"id"`
	CreatedOn	time.Time	`gorm:"type:timestamptz;not null;default:now()" json:"created_on"`
	Status		string		`gorm:"type:usage_status;not null;default:'success'" json:"status"`
	Key		uuid.UUID		`gorm:"type:uuid;not null;" json:"key"`
}


func MigrateUsage ( db *gorm.DB) {
	err := db.AutoMigrate(&Usage{})

	if err != nil {
		log.Fatalf("An error occured %s", err)
	}
}