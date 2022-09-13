package main

import (
	"fmt"
)

func main() {
	x := []int{4, 56, 89, 74, 5, 10}

	for i := 0; i < len(x); i++ {

		fmt.Printf("La valeur d'indice %d est %d\n", i, x[i])

	}

}
