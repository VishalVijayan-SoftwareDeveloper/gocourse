package main

import "fmt"

func main() {

	sequence := adder()
	fmt.Println(sequence()) //1
	fmt.Println(sequence()) //2
}
func adder() func() int {

	i := 0
	fmt.Println("previous value:", i)
	return func() int {
		i++
		fmt.Println("1 added to i", i)
		return i
	}
}
