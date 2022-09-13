package main

import (
	"fmt"
)

func main() {
	x := []int{}

	x = append(x, 2, 3, 8, 52, 10)
	y := []int{}
	y = append(y, x[1:]...)

	for _, v := range x {

		fmt.Println(v)

	}
	fmt.Printf("%T", x)
	fmt.Println(x[:])
	fmt.Println(x[:2])
	fmt.Println(x[3:5])
	for _, u := range y {

		fmt.Println(u)

	}
	fmt.Println(y)

}
