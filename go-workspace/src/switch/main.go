package main

import "fmt"

func main() {
	a := "Joao"

	switch a {
	case "Bob":
		fmt.Println("Hello Bob")
	case "Maria":
		fmt.Println("Hello Maria")
	case "Leandro":
		fmt.Println("Hello Leandro")
	default:
		fmt.Println("I did not find your name!")
	}
}
