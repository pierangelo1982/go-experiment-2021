package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	fmt.Println("ok")
	m := Message{"Pierangelo", "ciao mondo", 1294706395881547000}

	fmt.Println(m)
	fmt.Println(m.Name)

	fmt.Println("==================================================\n\tmarshal\t\n==================================================")

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)

	fmt.Println("==================================================\n\tunmarshal\t\n==================================================")

	var msg Message
	err = json.Unmarshal(b, &msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(msg)
	

}
