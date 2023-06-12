package main

import (
	"fmt"
)

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello")
}

func saysomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
	}
	saysomething(&p1)

	//fmt.Printf("I am happy to see %v say to the camera %v\n", p1.first, va)

}
