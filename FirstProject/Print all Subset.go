package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
   Go Program for
   Find all distinct subsets of a given set
*/

func allDistinctSubset(arr []int, n int) {
	// This is used to collect unique subset
	var record = make(map[string]bool)
	// Get value of 2ⁿ
	var size int = int(math.Pow(2, float64(n)))
	var subset string = ""
	// Outer loop
	for i := 0; i < size; i++ {
		// Inner loop are executing from 0...n-1
		for j := 0; j < n; j++ {
			if (i & (1 << j)) != 0 {
				subset += strconv.Itoa(arr[j]) + " "
				// Add subset
				record[subset] = true
			}
		}
		subset = ""
	}
	// Display distinct subsets
	for v := range record {
		fmt.Print("\n  ", v)
	}

}
func main() {

	var arr = []int{1, 2, 3}
	// Get the size of array
	var n int = len(arr)
	// Test
	allDistinctSubset(arr, n)
}
