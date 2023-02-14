package models

import (
	"geek-stash/models/extensions"
	"geek-stash/models/functions"
	"geek-stash/models/setup"

	"gorm.io/gorm"
)


func RunAllMigrations(db *gorm.DB){
	//Setup
	setup.Setup(db)
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
	//Keys
	MigrateKeys(db)
	//Usage
	MigrateUsage(db)
	//Functions
	functions.MigrateFunctions(db)
	
}