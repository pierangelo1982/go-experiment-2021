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
	//fmt.Println(driver)
	db := config.DatabaseConnection()
	rows, err := db.Query(`SELECT "id", "firstName", "lastName", "carNumber" FROM "Driver"`)
	if err != nil {
		panic(err)
	}

	var d []models.Driver
	for rows.Next() {
		err = rows.Scan(&driver.Id, &driver.FirstName, &driver.LastName, &driver.CarNumber)
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
