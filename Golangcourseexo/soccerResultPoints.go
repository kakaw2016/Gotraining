package main

import "fmt"

func teamPoint(status ...string) int {
	var totalPoint int
	for _, s := range status {
		switch {
		case s == "w":
			totalPoint += 3
		case s == "d":
			totalPoint += 1
		case s == "l":
			totalPoint += 0
		}

	}

	return totalPoint

}

func main() {
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	var score string
	fmt.Println("Enter the last match score, w for win, l for lost, d for draw")
	fmt.Scanln(&score)
	results = append(results, score)
	x := teamPoint(results...)
	fmt.Println("The final results are", x)

}
