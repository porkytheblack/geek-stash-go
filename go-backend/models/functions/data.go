package functions

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

type funcDef struct {
	name	string
	parameters	[]string
	declaration	string
	body	string
	returns	string
}


func (def *funcDef) genFunction() string {
	
	parameters := (func()string{
		if len(def.parameters) == 0 {
			return ""
		}else {
			return strings.Join(def.parameters, ", ")
		}
	})()
	str := fmt.Sprintf(`
		create or replace function %s ( %s )
		returns %s as
		$$
			%s ;
			begin
				%s
			end;
		$$ language plpgsql;
	`, def.name, parameters, def.returns, def.declaration, def.body )
	return str
}


func (def *funcDef) runFunc (db *gorm.DB, str string) {
	err := db.Exec(str).Error

	if err != nil {
		log.Fatalf("An error occured while working on function %s", err)
	}

}

func NewFuncMigrate (db *gorm.DB) {

	for _, def := range []funcDef{
		{
			name: "get_places",
			returns: "public.places[]",
			declaration: `
				declare a public.places[]
			`,
			body: `
				select * into a
				from public.places;
				return a;
			`,
		},
	} {
		d := def.genFunction()
		def.runFunc(db, d)
	}

}