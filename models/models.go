package models

import "gorm.io/gorm"


func RunAllMigrations(db *gorm.DB){

	//Profile
	MigrateProfile(db)
	//Franchise
	MigrateFranchise(db)
	//Gadgets
	MigrateGadgets(db)
	//Place
	MigratePlace(db)
	//Species
	MigrateSpecie(db)
	//Characters
	MigrateCharacter(db)
	
}