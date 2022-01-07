package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"webdemo/config"

	"webdemo/models"

	_ "github.com/lib/pq"
)

var driver struct{ models.Driver }

func HomePage(w http.ResponseWriter, req *http.Request) {
	db := config.DatabaseConnection()
	rows, err := db.Query(`SELECT d."id", d."firstName", d."lastName", d."carNumber", json_agg(json_build_object('name', country)) AS owner FROM "Driver" AS d LEFT JOIN "Team" AS t 
    ON d."teamId" = t."id" GROUP BY d."id";`)
	if err != nil {
		panic(err)
	}

	var d []models.Driver
	for rows.Next() {
		err = rows.Scan(&driver.Id, &driver.FirstName, &driver.LastName, &driver.CarNumber, &driver.Owner)
		if err != nil {
			panic(err)
		}
		fmt.Println(driver.FirstName, driver.LastName)
		d = append(d, driver.Driver)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(d)
}
