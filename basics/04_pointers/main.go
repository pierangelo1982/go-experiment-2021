package main

import (
	"fmt"
)

func main() {

	number1, number2 := 22, 88
	fmt.Println("prova pointers:")

	p1 := &number1
	p2 := &number2

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%x = %d \n", p1, *p1)
	fmt.Printf("%x = %d \n", p2, *p2)

	fmt.Println("======================================")

	var name string = "pierangelo"
	fmt.Println(name)

	var n *string = &name
	fmt.Println(n)
	fmt.Println(*n)
	fmt.Printf("%x = %s \n", n, *n)

	fmt.Println("======================================")

	fmt.Println(&number1)
	fmt.Println(&name)

}
