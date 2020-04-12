package main

import "fmt"

const (
	aa string = "Oiii" // não está visivel
	Bb int    = 66     // está visivel
	cc int    = 77     // não está visivel
)

const xvz int = 1333

// declarando uma variavel fora da função
var b int = 22

var c, d string = "Hello", "World"

func main() {
	// declarando uma variavel dentro da função
	a := 10
	b := "Hello"
	c := 10.56
	d := 'W'
	e := `Uouuuu`
	f := false
	const nome = "Leandro"

	// print o valor da variavel, por causa %v
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", nome)

	// print o tipo da variavel, por causa %T
	fmt.Printf("%T \n", a)    // int
	fmt.Printf("%T \n", b)    // string
	fmt.Printf("%T \n", c)    // float64
	fmt.Printf("%T \n", d)    // int32
	fmt.Printf("%T \n", e)    // string
	fmt.Printf("%T \n", f)    // bool
	fmt.Printf("%T \n", nome) // string
}
