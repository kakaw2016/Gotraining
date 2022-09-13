package main

import (
	"fmt"
)

type person struct {
	firstname string
}

func changename(perk *person) {
	perk.firstname = "ADISSIN"
}

func main() {
	p1 := person{
		firstname: "Tohouindo",
	}
	fmt.Println("Initial value of the person", p1)
	changename(&p1)
	fmt.Println("Changed value of the person", p1)
}
