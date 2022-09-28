package main

import (
	"fmt"
	"math"
	"sort"
)

func xpi(n int) (pi int, ok bool) {
	// If n is smaller than or equal to the largest cached prime,
	// we have an exact count
	var primes []int
	if i := sort.SearchInts(primes, n); i < len(primes) {
		if n == primes[i] {
			// n is the prime at index i
			pi = i + 1
		} else {
			// n is not prime and primes[j] < n for all j in [0,i)
			pi = i
		}
		ok = true
	} else {
		// n is larger than the largest prime in the cache;
		// use the estimate
		pi = int(float64(n) / (math.Log(float64(n)) - 1))
	}
	return
}

func sieve(n int) []int {
	switch {
	case n < 2:
		return []int{}
	case n == 2:
		return []int{2}
	}
	// a[i] == false ==> p=2*i+3 is a candidate prime
	// p in [3,n] ==> i in [0,(n-3)/2]
	length := 1 + (n-3)/2
	a := make([]bool, length, length)
	// Start with number 3 and consider only odd numbers
	sqrtn := int(math.Sqrt(float64(n)))
	for i, p := 0, 3; p <= sqrtn; p += 2 {
		if !a[i] {
			// 2*i+1 is a prime number; mark off its multiples
			for j := (p*p - 3) / 2; j < length; j += p {
				a[j] = true
			}
		}
		i++
	}
	// ps will store the computed primes; its initial capacity is based
	// an estimate of the prime-counting function pi(n)
	pi, _ := xpi(n)
	ps := make([]int, 1, pi)
	ps[0] = 2
	for i := 0; i < length; i++ {
		if !a[i] {
			ps = append(ps, 2*i+3)
		}
	}
	//fmt.Printf("Prime numbers generate are: %v\n", ps)
	return ps
}

func primebasedividingnumber(n int) (uint64, []uint64) {
	var (
		listofPrime               []int
		idealcombination          []uint64
		productcollection         []uint64
		positiveproductcollection []uint64
		productcollection2        []uint64
		productcollection3        []uint64
		divisibleprime            []uint64
		testednumber              uint64

		productcollection4 []uint64
		productcollection5 []uint64
	)
	listofPrime = sieve(n)

	for index := 0; index < len(listofPrime)-1; {

		initial := listofPrime[index]

		var indextwo int

		for indextwo < len(listofPrime)-1 {

			if initial != listofPrime[indextwo] {

				poweredvalue := uint64(math.Pow(float64(initial), float64(listofPrime[indextwo])))

				productcollection = append(productcollection, poweredvalue)

			}

			indextwo++

		}

		index++

		jindex := 0
		initial2 := listofPrime[jindex]

		if initial != initial2 {

			for jindex := 0; jindex < index; {

				poweredvalue := uint64(math.Pow(float64(initial), float64(initial2)))

				productcollection2 = append(productcollection2, poweredvalue)

				jindex++

			}
			//fmt.Println("FirstCC1", productcollection2)

		}

		//fmt.Println("FirstCd2", productcollection)
	}

	for index := len(listofPrime) - 1; index > 0; {

		initial := listofPrime[index]

		var indexthree int

		for indexthree < len(listofPrime)-1 {

			if initial != listofPrime[indexthree+1] {

				poweredvalue := uint64(math.Pow(float64(initial), float64(listofPrime[indexthree+1])))

				productcollection3 = append(productcollection3, poweredvalue)

			}

			indexthree++
		}

		jindex := 0
		initial2 := listofPrime[jindex]

		if initial != initial2 {

			for jindex := 0; jindex < index; {

				poweredvalue := uint64(math.Pow(float64(initial), float64(initial2)))

				productcollection4 = append(productcollection4, poweredvalue)

				jindex++

			}
			//fmt.Println("FirstCe1", productcollection4)

		}

		index--

		//fmt.Println("FirstCe2", productcollection3)
	}

	productcollection5 = append(productcollection, productcollection2...)

	productcollection6 := append(productcollection3, productcollection4...)

	productcollection7 := append(productcollection5, productcollection6...)

	sort.Slice(productcollection7, func(i, j int) bool { return productcollection7[i] < productcollection7[j] })

	//fmt.Println("FirstCC2", productcollection7)

	for i := 0; i < len(productcollection7)-1; i++ {

		if productcollection7[i] != productcollection7[i+1] && productcollection7[i] <= uint64(n) {

			positiveproductcollection = append(positiveproductcollection, productcollection7[i])

		}

	}
	//fmt.Println("Positive product collection", positiveproductcollection)

	for i := 0; i < len(positiveproductcollection)-1; i++ {

		if uint64(n)%positiveproductcollection[i] == 0 {
			divisibleprime = append(divisibleprime, positiveproductcollection[i])
		}

	}
	//fmt.Println("divisibleprime ", divisibleprime)

	for single := 0; single < len(divisibleprime)-1; {

		initial := divisibleprime[single]

		var singleone int

		for singleone < len(divisibleprime) {

			if initial*divisibleprime[singleone] == uint64(n) {

				idealcombination = append(idealcombination, initial, divisibleprime[singleone])
			}

			singleone++

		}

		single++

	}
	//fmt.Println("Idealcombination ", idealcombination)

	//fmt.Println("divisibleprime ", divisibleprime)
	testednumber = uint64(n)

	return testednumber, idealcombination
}

func main() {

	examplevalue := uint64(math.Pow(float64(800), float64(1)))

	chosenvalue, divisiblevalue := primebasedividingnumber(int(examplevalue))

	fmt.Printf("Testing value %v ;Divisible prime number of N is %v\n", chosenvalue, divisiblevalue)

}
