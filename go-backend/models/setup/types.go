package setup

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type tDef struct {
	ntype	string
	definition	string
}

func NewType (type_name string, definition string) string {
	str := fmt.Sprintf(`
	DO $$
	BEGIN
	   IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = '%s') THEN
		  CREATE TYPE %s as %s;
	   END IF;
	END $$;
	`, type_name, type_name, definition)

	return str
}

// Will have to add other related types later
func RegisterUserTypes (db *gorm.DB) {

	for _, t_def := range []tDef{
		{
			ntype: "key_status",
			definition: "enum('active', 'inactive')",
		},
		{
			ntype: "usage_status",
			definition: "enum('success', 'error')",
		},
	} {
		err := db.Exec(NewType(t_def.ntype, t_def.definition)).Error

		if err != nil {
			log.Fatalf("Unable to create type:: %s \n Error ::: %s", t_def.ntype, err)
		}
	}

}