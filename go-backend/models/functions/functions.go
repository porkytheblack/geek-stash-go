package functions

import "gorm.io/gorm"


func MigrateFunctions (db *gorm.DB) {
	NewFuncMigrate(db)
	SessionLogin(db)
	SessionLogout(db)
	GetSessionUID(db)
}