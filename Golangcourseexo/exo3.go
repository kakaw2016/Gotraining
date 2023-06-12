package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprint("firstvalue: ", x, " secondvalue: ", y, " thirdvalue: ", z)
	fmt.Printf("type %T\n", s)
	fmt.Println(s)
	s1 := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Printf(s1)
}
