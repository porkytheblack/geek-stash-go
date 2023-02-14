package models

import (
	"fmt"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type Keys struct {
	ID		uuid.UUID		`gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4(); unique" json:"id"`
	Key		string			`gorm:"type:text;not null" json:"key"`
	CreatedOn	time.Time	`gorm:"type:timestamptz;default:now()" json:"created_on"`
	UpdatedOn	time.Time	`gorm:"type:timestamptz;default:now()" json:"updated_on"`
	Status		string		`gorm:"type:key_status;default:'active'" json:"status"`
	Owner		uuid.UUID	`gorm:"type:uuid;not null;"`
	Usage	[]Usage			`gorm:"type:uuid;foreignKey:Key;references:ID"`
}

func MigrateKeys(db *gorm.DB) {
	err := db.AutoMigrate(&Keys{})

	if err != nil {
		log.Fatalf("Unable to migrate keys, encountered an error: %s", err)
	}

	fmt.Printf("Migrated Keys successfully!")

}