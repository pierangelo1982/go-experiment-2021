package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "0.0.0.0"
	port     = 5432
	user     = "postgres"
	password = "password1234"
	dbname   = "formulaonedb"
)

func DatabaseConnection() *sql.DB {
	// connection string
	//psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	// close database
	/* defer db.Close() */

	return db
}
