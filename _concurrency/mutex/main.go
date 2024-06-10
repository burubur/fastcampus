package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex
var stok int = 100

func beliProduk(seq int, jumlah int) {
	mutex.Lock()
	if stok >= jumlah {
		stok -= jumlah
		fmt.Printf("pembelian ke: %d berhasil dengan jumlah: %d\n", seq, jumlah)
	} else {
		fmt.Printf("pembelian ke: %d gagal, stok tidak cukup\n", seq)
	}
	mutex.Unlock()
}

func main() {
	for i := 0; i < 10; i++ {
		go beliProduk(i, 20)
	}

	time.Sleep(3 * time.Second)
}
