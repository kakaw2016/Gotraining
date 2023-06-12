package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	fmt.Println("La valeur d'indice est terrible")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("La valeur d'indice est faible.")
		wg.Done()

	}()

	go func() {
		fmt.Println("La valeur d'indice est grande")
		wg.Done()
	}()

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())

}
