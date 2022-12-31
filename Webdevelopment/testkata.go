package main

import (
	"fmt"
)

func Number(busStops [][2]int) int {

	var entrant, sortant, lastStation int

	for i := 0; i < len(busStops); i++ {

		entrant += busStops[i][0]

		sortant += busStops[i][1]
	}
	fmt.Println("First", entrant)
	fmt.Println("Second", sortant)

	lastStation = entrant - sortant

	return lastStation // Good Luck!
}

func main() {
	//arr1 := [][2]int{{10, 0}, {3, 5}, {5, 8}}

	arr2 := [][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}}

	//arr3 := [][2]int{{3, 0}, {9, 1}, {4, 8}, {12, 2}, {6, 1}, {7, 8}}

	//arr4 := [][2]int{{0, 0}}

	y := Number(arr2)
	fmt.Print(y)
}
