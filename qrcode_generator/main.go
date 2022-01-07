package main

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func main() {
	var name string
	fmt.Println("Inserisci il nome del file: ")
	fmt.Scanf("%s", &name)

	fmt.Println("inserisci l'url da codificare in qrcode: ")
	var urlToCOnvert string
	fmt.Scanf("%s", &urlToCOnvert)

	err := qrcode.WriteFile(urlToCOnvert, qrcode.Medium, 256, name+".png")
	if err != nil {
		fmt.Printf("Sorry couldn't create qrcode:,%v", err)
	}
}
