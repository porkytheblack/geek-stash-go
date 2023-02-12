package models

import (
	"geek-stash/models/extensions"
	"geek-stash/models/functions"

	"gorm.io/gorm"
)


func RunAllMigrations(db *gorm.DB){
	//Migrate Extensions
	extensions.MigrateExtensions(db)
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
	//Functions
	functions.MigrateFunctions(db)
	
}