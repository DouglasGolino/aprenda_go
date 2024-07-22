package main

import "fmt"

//var permite criar uma váriavel em package level scope
var x = 10

func main() {

	fmt.Print("Valor de X é ", x)

	QualquerCoisa(x)
}

func QualquerCoisa(n1 int) {
	fmt.Print("\r\nValor de N1 é ", n1)
}
