package main

import "fmt"

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
	var k int
	fmt.Print("\nEnter the k elements to sum up: ")
	fmt.Scan(&k)
	add := 0
	su := all[:k]
	for _, jsu := range su {
		add += jsu
		if k%2 != 0 && jsu%2 == 0 {
			fmt.Println("Totalmax  ODD pair equals to :", add)
		} else if k%2 != 0 && jsu%2 != 0 {
			fmt.Println("Totalmax  ODD pair equals to :", add)
		} else {
			fmt.Println("value is going to be Impair EVEN")
		}
	}
}
