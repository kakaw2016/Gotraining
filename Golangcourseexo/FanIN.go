package main

import (
	"fmt"
)

func send(e, o chan<- int) {
	for i := 1; i < 10; i++ {

		if i%2 == 0 {
			e <- i

		} else {
			o <- i

		}

		close(e)
		close(o)
	}

}

func main() {
	e := make(chan int)
	o := make(chan int)
	f := make(chan int)

	go send(e, o)

	receive1(e, f)

	receive2(o, f)

	for v := range f {
		fmt.Println(v)
	}

	fmt.Println("About to exit")

}

func receive1(e <-chan int, f chan<- int) {
	for v := range e {
		f <- v
	}
	close(f)

}
func receive2(o <-chan int, f chan<- int) {
	for v := range o {
		f <- v
	}
	close(f)

}
