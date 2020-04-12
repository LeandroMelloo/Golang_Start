package main

import "fmt"

var Y int = 33 // está sendo declarado globalmente, portanto pode ser utilizado fora do escopo de uma função como dentro do escopo da função.

func main() {
	x := 10 // está sendo declarado dentro do escopo da função, portanto vai ser utilizado dentro do escopo da função.
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	printZ()
}

func PrintY() {
	fmt.Println(y)
}
