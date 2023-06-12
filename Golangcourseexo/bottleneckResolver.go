package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Start a timer
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(3)

	// Perform an operation that may have a bottleneck
	go func() {
		for i := 0; i < 1000000; i++ {
			if i%52647 == 0 {
				fmt.Println("You did good", i)
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			if i%2647 == 0 {
				fmt.Println("You did good 2", i)
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			if i%62647 == 0 {
				fmt.Println("You did good 3", i)
			}
		}
		wg.Done()
	}()

	// Stop the timer
	elapsed := time.Since(start)

	// Print the elapsed time
	fmt.Println("Elapsed time:", elapsed)

	// Get the number of CPU cores
	numCPU := runtime.NumCPU()

	// Print the number of CPU cores
	fmt.Println("Number of CPU cores:", numCPU)

	// Get the current number of goroutines
	numGoRoutines := runtime.NumGoroutine()

	// Print the number of goroutines
	fmt.Println("Number of goroutines:", numGoRoutines)

	wg.Wait()

	// Get the current memory statistics
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	// Print the memory statistics
	fmt.Println("Alloc:", mem.Alloc)
	fmt.Println("TotalAlloc:", mem.TotalAlloc)
	fmt.Println("Sys:", mem.Sys)
}
