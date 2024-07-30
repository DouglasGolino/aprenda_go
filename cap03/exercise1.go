package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "James Bond"
	z := true

	//Um único comando
	fmt.Printf("%v,%v,%v\n", x, y, z)

	//Multiplos comandos
	fmt.Print(x, ",")
	fmt.Print(y, ",")
	fmt.Print(z)
}
