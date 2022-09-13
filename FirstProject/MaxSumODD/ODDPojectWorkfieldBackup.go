package main

import (
	"fmt"
)

func ReadN(all []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := Scan(&all[i]); m != 1 {
		panic(err)
	}
	ReadN(all, i+1, n-1)
}

func Scan(a *int) (int, error) {
	return fmt.Scan(a)
}

type aall []int

func (list aall) Has(a int) bool {
	for _, b := range list {
		if b == a {
			return true
		}

	}

	return false
}

func main() {
	var item int = 0
	var flag int = 0
	fmt.Println(`Enter the size of the Ensemble`)
	var n int
	if m, err := Scan(&n); m != 1 {
		panic(err)
	}
	fmt.Println(`Enter the integers`)
	all := make([]int, n)
	ReadN(all, 0, n)
	//fmt.Println(all)
	for i := 0; i <= n-1; i++ {
		item = all[i]
		flag = 0
		for j := i - 1; j >= 0 && flag != 1; {
			if item > all[j] {
				all[j+1] = all[j]
				j = j - 1
				all[j+1] = item
			} else {
				flag = 1
			}
		}
	}

	fmt.Printf("Array after sorting: \n")
	for i := 0; i <= n-1; i++ {
		fmt.Printf("%d ", all[i])
	}

	var ipair []int
	var impair int
	list := all[:]
	list = append(list, all...)

endone:
	for i := 0; i < len(list); i++ {
		if list[i]%2 != 0 {
			impair = list[i]
			ipair = append(ipair, impair)

			if len(ipair) == len(list) {
				fmt.Println("ALL are impair")
				break endone
			}

		}

	}

	//if len(ipair) != len(testList) {
	//fmt.Print("Contained both value\n")
	//
	//}
	var k int
	fmt.Print("\nEnter the k elements to sum up: ")
	fmt.Scan(&k)
	add := 0
	su := all[:k]
	jall := all[k-1:]
	var max int
	var cool int
	for _, jsu := range su {
		add += jsu
	}
looper:
	for i, j := 0, 0; i <= len(all)-1; i, j = i+1, j+1 {
		if (k%2 == 0 && all[i]%2 == 0 && add%2 == 0) || (k%2 != 0 && all[i]%2 == 0 && add%2 == 0) {
			fmt.Println("Total #1.2 equals to :", add)
			break looper
		}
		if k%2 == 0 && all[i]%2 != 0 && add%2 == 0 {
			fmt.Println("Total #1.1 equals to :", add)
			break looper
		}
		for _, jnub := range jall {
			switch {
			case all[k-1]%2 == 0:
				if jnub%2 != 0 && jnub >= max && add%2 != 0 {
					max = jnub
					cool = add - all[k-1] + max
					fmt.Println("Total #2 equals to :", cool)
					//fmt.Println("Total add equals to :", add)
					//fmt.Println("Total max equals to :", max)
					//fmt.Println("Total all[k-1] equals to :", all[k-1])
					//fmt.Println("Total SpeG2 equals to :", add)
					break looper
				}
			case all[k-1]%2 != 0:
				if jnub%2 == 0 && jnub >= max && add%2 != 0 {
					max = jnub
					cool = add - all[k-1] + max
					fmt.Println("Total #3 equals to :", cool)
					//fmt.Println("Total add equals to :", add)
					//fmt.Println("Total max equals to :", max)
					//fmt.Println("Total all[k-1] equals to :", all[k-1])
					break looper
				}
			default:
				fmt.Println("value is always going to be even")
			}
		}

	}

}
