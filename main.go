package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Brayan")

	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	fmt.Print(input)
}
