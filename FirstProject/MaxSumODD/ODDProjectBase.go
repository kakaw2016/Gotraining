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

	var ipair []int
	var impair int
	list := all[:]
	list = append(list, all...)
	var j, wSum, cool int = 0, 0, 0
	max := 0
looper:
	for i := 0; i < len(list); i++ {
		if k%2 != 0 && list[i]%2 != 0 {
			impair = list[i]
			ipair = append(ipair, impair)

			if len(ipair) == len(list) {

				fmt.Println("ALL SUM are impair")
				break looper
			}

		} else {
			for i := 0; i < len(list); i++ {

				for j < k {
					wSum += list[j]
					if wSum%2 == 0 {
						cool = wSum
						if cool > max {
							max = cool
						}
					}
					j++

				}

				for j < len(list) {
					wSum = wSum + list[j] - list[i]
					if wSum%2 == 0 {
						cool = wSum
						if wSum > max {
							max = cool
						}
					}
					i++
					j++
				}

				fmt.Println("Max SUM Equals to :", max)
				break looper

			}

		}
	}
}
