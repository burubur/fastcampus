package channel

import (
	"fmt"
	"time"
)

func Demo() {
	fmt.Println("program dijalankan...")
	// flow:
	// buatkan channel - done
	// buatkan penerima data lewat channel, async - done
	// buatkan pengirim data lewat channel, async - done
	// tunggu sampai program selesai - done

	messageCH := make(chan string)

	go func() {
		for i := 1; i <= 5; i++ {
			messageData := <-messageCH
			fmt.Println(messageData)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <= 5; i++ {
			messageCH <- "Ini adalah pesan dari goroutine."
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("program selesai")
}
