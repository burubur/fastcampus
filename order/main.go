package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// handle order
	order := Order{
		ID: 1,
	}

	// validate
	// 	distribute validation task -> fanOut CH
	fanOutCH := make(chan func(), 4)

	// 	wait all validation task -> fanIn CH
	fanInCH := make(chan bool, 4)

	var wg sync.WaitGroup

	// fanOut: distribute validation task
	wg.Add(4)
	go func() {
		fanOutCH <- func() { validatePaymentStatus(&order, &wg, fanInCH) }
		fanOutCH <- func() { validateSellerState(&order, &wg, fanInCH) }
		fanOutCH <- func() { validateStock(&order, &wg, fanInCH) }
		fanOutCH <- func() { validateShippingAddress(&order, &wg, fanInCH) }
		close(fanOutCH)
	}()

	// fanIn: collect result from all validation worker
	go func() {
		wg.Wait()
		close(fanInCH)
	}()

	// start worker
	for f := range fanOutCH {
		go f()
	}

	order.IsValid = true
	for result := range fanInCH {
		if !result {
			order.IsValid = false
		}
	}

	if !order.IsValid {
		fmt.Println("order is not valid, can't be processed.")
	} else {
		fmt.Println("order is valid")
	}
}

type Order struct {
	ID      int
	IsValid bool
}

func validatePaymentStatus(order *Order, wg *sync.WaitGroup, fanInCH chan<- bool) {
	defer wg.Done()

	time.Sleep(200 * time.Millisecond)
	fmt.Println("payment status validated: passed")
	fanInCH <- true
}

func validateSellerState(order *Order, wg *sync.WaitGroup, fanInCH chan<- bool) {
	defer wg.Done()

	time.Sleep(150 * time.Millisecond)
	fmt.Println("seller state validated: passed")
	fanInCH <- true
}

func validateStock(order *Order, wg *sync.WaitGroup, fanInCH chan<- bool) {
	defer wg.Done()

	time.Sleep(300 * time.Millisecond)
	fmt.Println("stock validated: passed")
	fanInCH <- true
}

func validateShippingAddress(order *Order, wg *sync.WaitGroup, fanInCH chan<- bool) {
	defer wg.Done()

	time.Sleep(50 * time.Millisecond)
	fmt.Println("shipping address validated: passed")
	fanInCH <- true
}
