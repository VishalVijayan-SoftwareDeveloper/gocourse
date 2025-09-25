package main

import "fmt"

func main() {
	//panic(interfacr{})

	check(10)
	check(-1)
}
func check(input int) {

	defer fmt.Println("Deffered statement 1")
	defer fmt.Println("Deffered statement 2")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	if input < 0 {
		panic("input must be non negative number")
	}
	fmt.Println("Valid input:", input)
}
