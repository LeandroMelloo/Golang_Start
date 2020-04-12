package main

import "fmt"

func main() {
	a := 6

	if a >= 10 {
		fmt.Println("A >= 10")
	} else if a > 5 {
		fmt.Println("A > 5")
	} else {
		fmt.Println("A é menor que 5")
	}

	b := false

	if x := "Cool"; b {
		fmt.Println(x)
		fmt.Println("True")
	} else {
		fmt.Println(x)
		fmt.Println("False")
	}
}
