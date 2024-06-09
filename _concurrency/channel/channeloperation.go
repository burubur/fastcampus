package channel

import (
	"fmt"
	"math/rand"
	"time"
)

func DemoChannel() {
	priceCh := make(chan StockPrice, 100)

	go priceProcessor(priceCh)

	// Start fetching stock prices for multiple symbols
	go fetchStockPrices(priceCh, "AAPL")
	go fetchStockPrices(priceCh, "GOOGL")
	go fetchStockPrices(priceCh, "MSFT")

	// Simulate some delay to allow fetching and processing
	time.Sleep(10 * time.Second)
}

type StockPrice struct {
	Symbol string
	Price  float64
	Time   time.Time
}

// this channel is to send only
func priceProcessor(priceCh <-chan StockPrice) {
	for price := range priceCh {
		fmt.Printf("Processing stock price: %s = %.2f at %s\n", price.Symbol, price.Price, price.Time.Format(time.RFC3339))
		// Simulate processing time
		time.Sleep(500 * time.Millisecond)
	}
}

// fetchStockPrices will be called by a scheduler
// this channel is to receive only
func fetchStockPrices(priceCh chan<- StockPrice, symbol string) {
	for {
		price := StockPrice{
			Symbol: symbol,
			Price:  rand.Float64() * 1000,
			Time:   time.Now(),
		}
		priceCh <- price
		// Simulate delay between price updates
		time.Sleep(time.Second)
	}
}
