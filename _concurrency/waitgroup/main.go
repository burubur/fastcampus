package main

import (
	"fmt"
	"sync"
)

func main() {
	FetchPricing()
}

func FetchPricing() {
	// flow:
	// start
	// fetch Indosat API
	// fetch Telkomsel API
	// fetch XL API
	// done

	var wg sync.WaitGroup

	var (
		indosatPricingResult   struct{}
		telkomselPricingResult struct{}
		xlPricingResult        struct{}
	)

	var (
		indosatErr   error
		telkomselErr error
		xlErr        error
	)

	wg.Add(1)
	go func() {
		indosatPricingResult, indosatErr = FetchIndosatAPI(&wg)
	}()

	wg.Add(1)
	go func() {
		telkomselPricingResult, telkomselErr = FetchTelkomselAPI(&wg)
	}()

	wg.Add(1)
	go func() {
		xlPricingResult, xlErr = FetchXLAPI(&wg)
	}()

	wg.Wait()
	_ = indosatPricingResult
	_ = telkomselPricingResult
	_ = xlPricingResult
	_ = indosatErr
	_ = telkomselErr
	_ = xlErr

	// cache all pricing result into Redis
	fmt.Println("fetched all pricings")
}

func FetchIndosatAPI(wg *sync.WaitGroup) (data struct{}, err error) {
	defer wg.Done()

	fmt.Println("fething pricing from Indosat")
	return data, nil
}

func FetchTelkomselAPI(wg *sync.WaitGroup) (data struct{}, err error) {
	defer wg.Done()

	fmt.Println("fething pricing from Telkomsel")
	return data, nil
}

func FetchXLAPI(wg *sync.WaitGroup) (data struct{}, err error) {
	defer wg.Done()

	fmt.Println("fething pricing from XL")
	return data, nil
}
