package main

import (
	"fmt"
	"time"
)

func measureExecution(handler func()) func() {

	return func() {
		start := time.Now()
		handler()
		fmt.Printf("Execution took %v\n", time.Since(start))
	}

}

func main() {

	createOrder := func() {
		fmt.Println("Creating new sales order")
		time.Sleep(2 * time.Second)
		fmt.Println("Order created successfully")
	}
	wrapped := measureExecution(createOrder)
	wrapped()

}
