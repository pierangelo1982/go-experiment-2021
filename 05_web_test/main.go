package main

import (
	"log"
	"net/http"
	"webdemo/routes"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", routes.HomePage)
	http.HandleFunc("/test", routes.TestPage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
