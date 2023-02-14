package setup

import "gorm.io/gorm"


func Setup ( db *gorm.DB ) {
	//Types
	RegisterUserTypes(db)
}