package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go printMessage("Hello from Goroutine 1", &wg)

	wg.Add(1)
	go printMessage("Hello from Goroutine 2", &wg)

	wg.Wait()
	fmt.Println("program selesai")
}

func printMessage(message string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}
