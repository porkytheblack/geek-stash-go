package functions

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)


func SessionLogin (db *gorm.DB) {
	err := db.Exec(fmt.Sprintf(`
	  create or replace function login (token text)
	  returns void as $$
	  declare claims json;
	  declare id text;
	  begin
		claims := row_to_json(jwt.verify(token, '%s'))->'payload';
		id := claims->'sub'::text;
		perform
			set_config('request.jwt.claim.sub', id::text, TRUE);
	  end;
	  $$ language plpgsql;
	`, os.Getenv("AUTH0_SECRET"))).Error
	if err != nil {
		log.Printf("An error occured creating session login function:: %s", err)
	}
}


func SessionLogout (db *gorm.DB) {
	err := db.Exec(`
		create or replace function logout ()
		returns void as $$
			begin
				perform
					set_config('request.jwt.claim.sub', NULL, TRUE);
			end;
		$$ language plpgsql;
	`).Error

	if err != nil {
		log.Printf("An error occured logging out of session:: %s", err)
	}
}


func GetSessionUID (db *gorm.DB) {
	err := db.Exec(`
		create or replace function uid ()
		returns uuid as
		$$
		declare a uuid;
		begin
			a := regexp_replace(nullif(current_setting('request.jwt.claim.sub', true), ''), '"', '', 'g')::uuid;
			return a;
		end;
		$$ language plpgsql;
	`).Error

	if err != nil {
		log.Printf("An error occured creating get uid function,:: %s", err)
	}

}