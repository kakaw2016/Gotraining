package main

import (
	"fmt"
)

type person struct {
	firstname               string
	lastname                string
	favoriteicecreamflavors []string
}

func main() {
	p1 := person{
		firstname:               "Cameroon",
		lastname:                "Johnson",
		favoriteicecreamflavors: []string{"Chocolate", "peach"},
	}
	p2 := person{
		firstname:               "Camery",
		lastname:                "John",
		favoriteicecreamflavors: []string{"Chocolate", "Bananas"},
	}
	p3 := person{
		firstname:               "Leroon",
		lastname:                "Annson",
		favoriteicecreamflavors: []string{"Chocolate", "Ananas"},
	}
	x := []person{}
	x = append(x, p1, p2, p3)
	fmt.Println(x)
	fmt.Println("Second Person Last name", p2.lastname)
	for i, v := range p2.favoriteicecreamflavors {
		fmt.Printf("P2 Favorite Flavor %v and %v\n", i, v)
	}

	for _, v := range x {
		fmt.Println(v)
	}
	data := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
		p3.lastname: p3,
	}
	fmt.Printf("Map containing structure person element \n%v\n", data)
	for _, v := range data {

		fmt.Println(v.firstname)
		fmt.Println(v.lastname)

		for i, val := range v.favoriteicecreamflavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}

	for _, v := range data {
		fmt.Println(p2.firstname, v.lastname)
	}

}
