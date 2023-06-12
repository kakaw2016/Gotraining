package main

import (
	"fmt"
)

type monamie int
type monfrere monamie

var x monamie
var y monfrere
var z int

func main() {
	fmt.Println(x)
	fmt.Printf("Type %T\n", x)
	x = 2000
	fmt.Println(x)
	y = monfrere(x)
	fmt.Printf("Type %T\n", y)
	fmt.Println(y)
	z = int(y)
	fmt.Println(z)

}
