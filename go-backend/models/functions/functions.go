package functions

import "gorm.io/gorm"


func MigrateFunctions (db *gorm.DB) {
	SessionLogin(db)
	SessionLogout(db)
	GetSessionUID(db)
}