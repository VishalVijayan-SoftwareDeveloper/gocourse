package main

import "fmt"

func invoiceProcessor() func(float64) float64 {

	total := 0.0

	return func(amount float64) float64 {
		total += amount
		return total
	}

}

func main() {
	process := invoiceProcessor()

	fmt.Printf("Invoice proccesed, runnning total:=%.2f\n", process(1000))
	fmt.Printf("Invoice proccesed, runnning total:=%.2f\n", process(100))
	fmt.Printf("Invoice proccesed, runnning total:=%.2f\n", process(500.50))
	fmt.Printf("Invoice proccesed, runnning total:=%.2f\n", process(1500.60))

}
