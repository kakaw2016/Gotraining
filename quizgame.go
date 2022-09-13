package main

import "fmt"

//func main() {
//fmt.Println("Welcome to my quiz game!")
//
//name := "Olivier"
//age := 21
//fmt.Println(name)
//fmt.Printf("Hello %v, Your age is %v!", name, age)
//}

func main() {
	fmt.Println("Welcome to my quiz game!")
	fmt.Printf("What is your name?\n")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Your name is %v.\n", name)
	fmt.Printf("Enter your age:\n")
	var age uint
	fmt.Scan(&age)
	if age >= 20 {
		fmt.Printf("Your age, %v year old, is right for the game.\n", age)
	} else {
		fmt.Println("you cannot play!")
		return
	}

	fmt.Printf("What is the better card? RTX 30800 or RTX 3090 ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 30800" || answer+" "+answer2 == "rtx 30800" {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect!")
	}
}
