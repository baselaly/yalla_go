package request

import (
	"fmt"
	"log"
	"strings"
	connection "yalla_go/db"

	"github.com/thedevsaddam/govalidator"
)

func init() {

	// unique field rule
	govalidator.AddCustomRule("unique", func(field string, rule string, message string, value interface{}) error {
		db := connection.Connect()
		err := db.Ping()

		if err != nil {
			log.Fatalln(err)
		}
		defer connection.Close(db)

		var count int
		val := value.(string)
		table := strings.TrimPrefix(rule, "unique:")

		rows, err := db.Query("SELECT COUNT(*) FROM "+table+" WHERE "+field+" = ?", val)
		if err != nil {
			log.Fatalln(err)
		}

		for rows.Next() {
			rows.Scan(&count)
		}

		if count > 0 {
			return fmt.Errorf("The %s already taken", field)
		}

		return nil
	})

	// confirmation rule
	govalidator.AddCustomRule("confirmed", func(field string, rule string, message string, value interface{}) error {
		val := value.(string)
		confirmation := strings.TrimPrefix(rule, "confirmed:")
		if val != confirmation {
			return fmt.Errorf("The %s field need confirmation", field)
		}
		return nil
	})
}
