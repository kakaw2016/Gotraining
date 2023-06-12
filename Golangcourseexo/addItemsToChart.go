package main

import "fmt"

type Cart struct {
	prices []float32
}

func (x *Cart) show() {
	var sum float32 = 0.0
	//calculate the sum of all prices in the Cart
	for _, v := range x.prices {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	c := Cart{}
	var n int
	var m float32
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		fmt.Scanln(&m)
		c.prices = append(c.prices, m)

	}
	// take n inputs and add them to the Cart

	//call the show() method of the Cart
	c.show()

}
