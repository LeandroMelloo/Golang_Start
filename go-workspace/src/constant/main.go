package main

import "fmt"

const (
	aa string = "Oiii" // não está visivel
	Bb int    = 66     // está visivel
	cc int    = 77     // não está visivel
)

const xvz int = 1333

func main() {
	// declarando uma variavel dentro da função
	const nome = "Leandro"

	// print o valor da variavel, por causa %v
	fmt.Printf("%v \n", nome)

	// print o tipo da variavel, por causa %T
	fmt.Printf("%T \n", nome) // string
}
