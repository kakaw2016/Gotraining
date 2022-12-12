package main

import "fmt"

func main() {
	menu := [6]string{"Water", "Burger", "Cake", "Soup", "Soda", "Fries"}

	//your code goes here
	var input int
	fmt.Println("Enter your preferred menu number")
	fmt.Scanln(&input)

	if input >= 0 && input < len(menu) {
		fmt.Println("Congratulations you will have", menu[input])

	} else {
		fmt.Println("Invalid choice")
	}

}
