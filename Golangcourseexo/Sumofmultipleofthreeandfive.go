//Cool exercise where we must find divisible of 3 then 5
//Add all the integer that filled the condition above
//Print the sum that is taken within a range of integers

package main

import (
	"fmt"
)

var checkvalue, goodmatch1, goodmatch2 int

func multipleofthree(checkvalue int) int {
	if checkvalue%3 == 0 {
		goodmatch1 += checkvalue

	}
	return goodmatch1

}

func multipleoffive(checkvalue int) int {

	if checkvalue%5 == 0 {
		goodmatch2 += checkvalue

	}
	return goodmatch2

}

// this step of the code defines a range of integer where the addition of divisible will be calculated
func sumofmultiple(checkvalue int, limit int) int {
	for ; checkvalue <= limit; checkvalue++ {

		multipleoffive(checkvalue)
		multipleofthree(checkvalue)
	}

	return multipleoffive(checkvalue) + multipleofthree(checkvalue)

}

func main() {
	var finalresult int

	finalresult = sumofmultiple(3000, 3040)

	fmt.Printf("The sum of the multiple of 3 and 5 is: %v\n", finalresult)

	fmt.Println("End of the program")
}
