package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)

	receive(eve, odd, quit)

}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the eve channel:", v)

		case v := <-o:
			fmt.Println("From the odd channel:", v)

		case i, ok := <-q:
			if !ok {
				fmt.Println("From the Quit channel OK:", i, ok)
				return
			} else {
				fmt.Println("From the Quit channel OK:", i)
				return

			}
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	close(q)

}
