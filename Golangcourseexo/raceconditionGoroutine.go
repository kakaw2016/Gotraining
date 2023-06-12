package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	incrementation := 0
	gs := 100

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		fmt.Println("MId value", incrementation)

		go func() {

			v := incrementation
			runtime.Gosched()
			v++
			incrementation = v

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value", incrementation)

}
