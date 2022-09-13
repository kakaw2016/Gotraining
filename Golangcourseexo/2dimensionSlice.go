package main

import (
	"fmt"
)

func main() {
	x := [][]string{{"James", "Bond", "Shaked, not stirred"}, {"Miss", "Moneyppenny", "Hellooooo, James."}}

	for _, v := range x {

		fmt.Println(v)

	}
	for i, v := range x {
		fmt.Println("record ", i)
		for _, p := range v {

			fmt.Println(p)
		}

	}
	fmt.Printf("%T", x)

}
