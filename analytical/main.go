package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// simulate fetching order
	// simulate filtering order
	// simulate analysing order item
	// simulate storing the analysis result to big query
	// fetch -> filter -> analyse -> store
	orderCH := make(chan Order)
	filterdOrderCH := make(chan Item)
	analysisReportCH := make(chan AnalysisReport)

	// run stage of pipeline concurrently
	go fetchOrder(orderCH)
	go filterOrder(orderCH, filterdOrderCH)
	go analyseOrder(filterdOrderCH, analysisReportCH)
	go storeAnalysisReport(analysisReportCH)

	select {}
}

type (
	Order struct {
		ID    int
		Items []Item
	}

	Item struct {
		ItemID       int
		Category     string // to be filtered, specific for "digital" category
		ProviderName string // telkomsel, indosat, esia
		Price        float64
	}

	AnalysisReport struct {
		Category     string
		AveragePrice float64
		MinPrice     float64
		MaxPrice     float64
	}
)

func fetchOrder(orderCH chan<- Order) {
	for i := 1; ; i++ {
		orderData := Order{
			ID: i,
			Items: []Item{
				{
					ItemID:       i*10 + 2,
					Category:     "digital",
					ProviderName: "telkomsel",
					Price:        rand.Float64() * 10000,
				},
			},
		}

		time.Sleep(500 * time.Millisecond)
		orderCH <- orderData
		fmt.Println("fetched the order")
	}
}

func filterOrder(orderCH <-chan Order, filterdOrderCH chan<- Item) {
	for order := range orderCH {
		for _, item := range order.Items {
			if item.Category == "digital" {
				time.Sleep(200 * time.Millisecond)
				filterdOrderCH <- item
				fmt.Println("filtered order detail")
			}
		}
	}
}

func analyseOrder(filterdOrderCH <-chan Item, analysisReportCH chan<- AnalysisReport) {
	for item := range filterdOrderCH {
		// do the analysis here
		result := AnalysisReport{
			Category:     item.Category,
			MinPrice:     item.Price,     // to simplify the calculation
			MaxPrice:     item.Price * 2, // to simplify the calculation
			AveragePrice: item.Price,     // to simplify the calculation
		}
		analysisReportCH <- result
		fmt.Println("analyzed the order detail")
	}
}

func storeAnalysisReport(analysisReportCH <-chan AnalysisReport) {
	for report := range analysisReportCH {
		// store to big query
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("stored the analysis report to BQ with result: %+v\n\n", report)
	}
}
