package main

import (
	"log"
	"net/http"
	"os"
	"pierangelo1982/httpserverdemo/handlers"
)

func main() {
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":8080", sm)
}
