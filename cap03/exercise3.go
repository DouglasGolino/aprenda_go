package main

import (
	"fmt"
)

var x int = 12
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\n%v\n%v", x, y, z)
	fmt.Printf("\n\nVariável escrita com Sprint:\n%v\n", s)
}
