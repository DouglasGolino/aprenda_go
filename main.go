package main

import "fmt"

func main() {
	const name, age = "Douglas", 40
	numerodebytes, erros := fmt.Println(name, "is", age, "years old.")
	fmt.Println(numerodebytes, erros)
}
