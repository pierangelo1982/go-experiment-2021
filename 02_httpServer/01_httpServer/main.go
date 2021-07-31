package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Tester struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
	Id   int64  `json:"id"`
}

func main() {
	// curl -v -d'pierangelo' localhost:8080/test
	http.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("hello test")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			//rw.WriteHeader(http.StatusBadRequest)
			//rw.Write([]byte("Oooops somethings goes wrong!"))
			// rw sopra posson essere sostituiti con questa riga
			http.Error(rw, "Ooops something goes wrong!", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s \n", d)
	})

	http.HandleFunc("hello", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
