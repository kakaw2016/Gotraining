package main

import (
	"fmt"
)

type monamie int

var x monamie

func main() {
	fmt.Println(x)
	fmt.Printf("Type %T\n", x)
	x = 42
	fmt.Println(x)

}
