package main

import "fmt"

func main() {
	foo()
	func() {
		fmt.Println("Anonymous func Ran")
	}()

	func(x int) {
		fmt.Println("The meaning of life", x)
	}(32)

}
func foo() {
	fmt.Println("foo ran")
}
