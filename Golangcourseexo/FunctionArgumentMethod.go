package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)

}

func test(ss secretAgent) {
	fmt.Println("I am testing", ss.first, ss.last)

}

func testt(sss secretAgent) string {
	return fmt.Sprintln("I am really testing", sss.first, sss.last)

}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	test(sa1)
	fmt.Print(testt(sa1))
}
