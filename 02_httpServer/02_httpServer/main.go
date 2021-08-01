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
	hg := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", hg)

	http.ListenAndServe(":8080", sm)
}

// curl -v -d"pierangelo" localhost:8080
// curl -v localhost:8080/goodbye
