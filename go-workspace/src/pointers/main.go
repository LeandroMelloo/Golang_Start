package main

import "fmt"

func xpto(a *int) int {
	*a = 100
	return *a
}

func main() {
	b := 10
	fmt.Println(xpto(&b)) // 100
	fmt.Println(b)        // 100, porque está mudando o endereçamento de memoria de "b"

	// guardando "x" em um espaço da memoria
	x := 10
	// &x => referencia o valor de memoria de "x"
	fmt.Println(&x) // 0xc0000120a0
	fmt.Println(x)  // 10

	// guardando "y" pro mesmo local de memoria que está referenciado minha variavel "x".
	y := &x
	// irá guardar o valor de "x" na variavel "y"
	fmt.Println(y) // 0xc0000120a0
	// desreferenciando "y"
	fmt.Println(*y) // 10

	*y = 20
	fmt.Println(y) // 0xc0000120a0
	fmt.Println(x) // 20

	var z *int = &x
	fmt.Println(z)  // 0xc0000120a0
	fmt.Println(*z) // 20
}
