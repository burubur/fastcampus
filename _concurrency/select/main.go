package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// isInterruptedCH
	// sigCH

	sigCH := make(chan os.Signal, 1)
	signal.Notify(sigCH, syscall.SIGTERM, syscall.SIGINT)

	interruptionCH := make(chan bool)
	anotherCH := make(chan any)

	go func() {
		for {
			select {
			case <-interruptionCH:
				fmt.Printf("Task goroutine diberhentikan ...")
				return
			case <-anotherCH:
				// do something here
			default:
				fmt.Println("menjalanankan tugas ...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	select {
	case interruptionSignal := <-sigCH:
		fmt.Printf("interruption signal trigerred: %v\n", interruptionSignal)
		fmt.Println("shutting down the program gracefully...")
		close(interruptionCH)
	}
}
