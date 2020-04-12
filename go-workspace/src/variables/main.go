package main

import "fmt"

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

	// print o valor da variavel, por causa %v
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)

	// print o tipo da variavel, por causa %T
	fmt.Printf("%T \n", a) // int
	fmt.Printf("%T \n", b) // string
	fmt.Printf("%T \n", c) // float64
	fmt.Printf("%T \n", d) // int32
	fmt.Printf("%T \n", e) // string
	fmt.Printf("%T \n", f) // bool
}
