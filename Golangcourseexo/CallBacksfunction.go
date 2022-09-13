package main

import (
	"fmt"
)

func main() {
	a := []int{2, 5, 4, 8, 7, 9}

	t1 := sum(a...)
	fmt.Println("All number sum: ", t1)
	s2 := even(sum, a...)
	fmt.Println("Even number sum", s2)
	s3 := odd(sum, a...)
	fmt.Println("Odd number sum", s3)

}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v

	}
	return total
}
func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 1 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
