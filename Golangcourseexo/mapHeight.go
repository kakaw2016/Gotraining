package main

import "fmt"

func main() {
	var sum, count, average float32
	team := map[string]float32{
		"P1": 1.98,
		"P2": 2.05,
		"P3": 1.89,
		"P4": 2.0,
		"P5": 2.11}
	for key, height := range team {
		sum += height
		if key != "" {
			count += 1
		}

	}
	average = sum / count
	fmt.Println(average)

}
