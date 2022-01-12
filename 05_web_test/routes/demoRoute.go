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
var team struct{ models.Team }

func HomePage(w http.ResponseWriter, req *http.Request) {
	db := config.DatabaseConnection()
	customQuery := `
		SELECT c.id, c."firstName", c."lastName", c."carNumber", c."teamId", json_agg(json_build_object(name, country)) AS owner FROM "Driver" AS c 
    	LEFT JOIN "Team" AS t ON t."id"=c."teamId" GROUP BY c."id"
	`
	rows, err := db.Query(customQuery)
	if err != nil {
		panic(err)
	}

	var d []models.Driver

	for rows.Next() {
		err = rows.Scan(&driver.Id, &driver.FirstName, &driver.LastName, &driver.CarNumber, &driver.TeamId, &driver.Owner)
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
