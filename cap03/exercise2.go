package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Printf("Valores padrão de int, string e bool \n")
	fmt.Printf("Valor padrão: %v para o tipo %T\n", x, x)
	fmt.Printf("Valor padrão: %v para o tipo %T\n", y, y)
	fmt.Printf("Valor padrão: %v para o tipo %T\n", z, z)
}
