package main

import (
	"fmt"
	"time"
)

func task1() {
	fmt.Println("Running task 1...")
}

func task2() {
	fmt.Println("Running task 2...")
}

func main() {
	ticker := time.NewTicker(time.Minute * 30)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				task1()
				task2()
				fmt.Println("Tasks ran at", t)
			}
		}
	}()

	time.Sleep(time.Minute * 90)
	ticker.Stop()
	done <- true
	fmt.Println("Tasks stopped.")
}
