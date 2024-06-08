package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Halo, dari main 1")
	go sayHello()
	fmt.Println("Halo, dari main 2")

	time.Sleep(time.Second * 5)
	fmt.Println("Halo, dari main 3")
}

func sayHello() {
	fmt.Println("Halo, dari GoRoutine")
}
