package main

import "fmt"

func main() {
	slice := make([]int, 2, 5)
	fmt.Println(slice)                           // [0, 0]
	fmt.Println(len(slice))                      // 2 => tamanho do slice
	fmt.Println(cap(slice))                      // 5 => capacidade do slice
	slice = append(slice, 1, 12, 33, 89, 77, 90) // 8 => tamanho dom slice = 2 + append = 6. Totalizando 8
	fmt.Println(slice)                           // [0, 0, 0, 0, 0, 1, 12, 33, 89, 77, 90]
	fmt.Println(len(slice))                      // 8 => passou do tamanho definido, porém a capacidade dobra
	fmt.Println(cap(slice))                      // 10 => dobrou a capacidade do array

	for i := 0; i < 20; i++ {
		slice = append(slice, i)
		fmt.Println("Len: ", len(slice), "Cap: ", cap(slice))
	}

	sliceTest := slice
	slice = append(slice, 1, 2, 3, 4)
	slice[0] = 10
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(sliceTest) //[0, 0]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	sliceString := []string{
		"Hello",
		"World",
		"Much",
		"Better",
		"Better2",
	}
	fmt.Println(sliceString[0])   // [Hello]
	fmt.Println(sliceString[:2])  // [Hello World]
	fmt.Println(sliceString[1:2]) // [World]
	fmt.Println(sliceString[2:4]) // [Much Better]
	fmt.Println(sliceString[2:])  // [Much Better Better2]
}
