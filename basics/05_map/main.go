package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")

	m := make(map[string]int)

	m["k1"] = 5
	m["k2"] = 10

	fmt.Println(m)

	v := m
	fmt.Println(v["k1"])

	fmt.Println("==============================================")

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}
}
