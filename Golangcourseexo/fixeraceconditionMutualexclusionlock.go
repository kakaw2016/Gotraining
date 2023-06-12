package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	incrementation := 0
	gs := 100

	wg.Add(gs)
	var m sync.Mutex
	for i := 0; i < gs; i++ {

		go func() {
			m.Lock()

			v := incrementation

			v++
			incrementation = v

			fmt.Println("Mid value", incrementation)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End value", incrementation)

}
