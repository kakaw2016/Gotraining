package main

import (
	"fmt"
)

func zero(z *int) {
	fmt.Println(z)
	*z = 0
}

func main() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Printf("Type of b %T and value %v\n", b, b)
	fmt.Println(*b)
	fmt.Printf("Type of *b %T and value %v\n", *b, *b)

	*b = 42
	fmt.Printf("Type of a %T and value %v\n", a, a)
	fmt.Printf("Type of *b %T and value %v\n", *b, *b)

	x := 5
	fmt.Println("Memory address of x", &x)
	zero(&x)
	fmt.Printf("Type of x %T and value %v\n", x, x)
}
