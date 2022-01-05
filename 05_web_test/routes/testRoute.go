package routes

import (
	"fmt"
	"net/http"

	"webdemo/config"

	_ "github.com/lib/pq"
)

func TestPage(w http.ResponseWriter, req *http.Request) {
	db := config.DatabaseConnection()
	rows, err := db.Query(`SELECT "id", "firstName", "lastName", "carNumber" FROM "Driver"`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		var carNumber int

		err = rows.Scan(&id, &firstName, &lastName, &carNumber)
		if err != nil {
			panic(err)
		}

		fmt.Println(firstName, lastName, carNumber)
	}

	fmt.Fprintf(w, "Hello World!")
}
