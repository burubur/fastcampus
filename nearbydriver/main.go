package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// cached on Redis
// var (
// 	TopDriverIDs      []uuid.UUID
// 	ReliableDriverIDs []uuid.UUID
// 	NormalDriverIDs   []uuid.UUID
// )

func main() {
	fmt.Println("start...")
	FindNearbyDriverIDs()
	fmt.Println("finished.")
}

func FindNearbyDriverIDs() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	doneCH := make(chan struct{}, 3)

	// findTopDriverIDs - optional
	go func() {
		findTopDriverIDs(ctx)
		doneCH <- struct{}{}
	}()

	// findReliableDriverIDs - optional
	go func() {
		findReliableDriverIDs(ctx)
		doneCH <- struct{}{}
	}()

	// findNormalDriverIDs - mandatory
	go func() {
		findNormalDriverIDs(ctx)
		doneCH <- struct{}{}
	}()

	for i := 0; i < 3; i++ {
		<-doneCH
	}

	fmt.Println("all goroutine finished")
}

func findTopDriverIDs(ctx context.Context) (driverIDs []uuid.UUID) {
	// simulate to have a long running query
	select {
	case <-time.After(1500 * time.Millisecond):
		fmt.Println("top drivers found")
		return []uuid.UUID{uuid.New(), uuid.New()}
	case <-ctx.Done():
		fmt.Println("find top drivers canceled")
		return []uuid.UUID{}
	}
}

func findReliableDriverIDs(ctx context.Context) (driverIDs []uuid.UUID) {
	// simulate to have a long running query
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("reliable drivers found")
		return []uuid.UUID{uuid.New(), uuid.New()}
	case <-ctx.Done():
		fmt.Println("find reliable drivers canceled")
		return []uuid.UUID{}
	}
}

func findNormalDriverIDs(ctx context.Context) (driverIDs []uuid.UUID) {
	fmt.Println("normal drivers found")
	return []uuid.UUID{uuid.New()}
}
