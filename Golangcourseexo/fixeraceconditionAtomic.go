package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incrementation int64
	gs := 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {

		go func() {
			atomic.AddInt64(&incrementation, 1)
			fmt.Println(atomic.LoadInt64(&incrementation))

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End value", incrementation)

}
