package main

import "fmt"

func main() {
	var x [10]int
	fmt.Println(x)      // [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
	fmt.Println(len(x)) // 10 => tamanho do array

	x[1] = 3
	x[5] = 6
	x[9] = 9
	fmt.Println(x)    // [0, 3, 0, 0, 0, 6, 0, 0, 0, 9]
	fmt.Println(x[0]) // 0

	var y [10]string // inicializei o array
	fmt.Println(y)   // [                   ]

	z := [5]int{1, 2, 3, 4, 5}
	fmt.Println(z) // [1, 2, 3, 4, 5]
}
