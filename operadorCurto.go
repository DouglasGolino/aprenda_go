package main

import (
	"fmt"
)

// operador curto := só funciona dentro de codeblock

func main() {

	x := 10

	//usando format print para especificar um valor de X em v e um tipo em T
	fmt.Printf("x: %v, %T\n", x, x)

	y := "Um valor qualquer"

	fmt.Printf("y: %v, %T\n\n", y, y)

	//neste caso é possível usar o operador de declaração em X porque estou criando uma varivável nova z
	// se não houver uma nova váriavel sendo criada eu somente poderia usar um operador de atribuição
	x, z := 15, 22

	fmt.Print("x: ", x, " z ", z)
}
