package main

import "fmt"

var Y int = 33 // está sendo declarado globalmente, portanto pode ser utilizado fora do escopo de uma função como dentro do escopo da função.

func PrintY() {
	fmt.Println(y)
}
