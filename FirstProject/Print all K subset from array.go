package main

import (
	"fmt"
)

//Print all subsequences in a given size k
func print_subsequences(collection []int, output []int, n, k, next, location int) {

	if location == k {
		// When get new subsequence of k size
		fmt.Print("[")
		for i := 0; i < k; i++ {
			fmt.Print(" ", collection[output[i]])
		}

		var sum int = 0

		for i := 0; i < k; i++ {
			sum = sum + collection[output[i]]

		}
		fmt.Println("] SUM", sum)

	} else {
		for i := next; i < n; i++ {
			// Get location of subset collection
			output[location] = i
			print_subsequences(collection, output, n, k, i+1, location+1)
		}

	}

}

// Handles the request to print subsequence element
func subsequences(collection []int, n, size int) {
	if n <= 0 || size <= 0 && size >= n {
		//if not a valid inputs
		return
	}
	//Used to display result
	output := make([]int, size)
	//fmt.Print("\n Subsequences of size \n", size)
	print_subsequences(collection, output, n, size, 0, 0)

}

func main() {

	arr := []int{1, 2, 3, 4, 5}
	n := len(arr)
	size := 2
	subsequences(arr, n, size)

}
