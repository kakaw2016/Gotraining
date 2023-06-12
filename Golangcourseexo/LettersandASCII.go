package main

import (
	"fmt"
)

func main() {
	for i := 85; i <= 90; i++ {
		fmt.Printf("%b\t %#U\n", i, i)
		fmt.Print("\n--------\n")
		fmt.Printf("%.*s\n", i, "*****")
	}

}
