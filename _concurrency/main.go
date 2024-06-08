package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printNumbers(1, &wg)

	wg.Add(1)
	go printNumbers(2, &wg)

	wg.Wait()
	fmt.Println("program selesai")
}

func printNumbers(jobID int, wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("job id: %d, menjalankan task %d\n", jobID, i)
	}

	wg.Done()
}
