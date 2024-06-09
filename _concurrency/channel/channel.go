package channel

import (
	"fmt"
	"time"
)

func Demo() {
	fmt.Println("program dijalankan...")

	messageCH := make(chan string, 2)

	go func() {
		for i := 1; i <= 12; i++ {
			fmt.Printf("data ke: %d dikirim\n", i)
			messageCH <- fmt.Sprintf("data dari goroutine ke: %d", i)
		}
	}()

	go func() {
		for {
			messageData := <-messageCH
			fmt.Printf("data diterima: %s\n", messageData)
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("program selesai")
}
