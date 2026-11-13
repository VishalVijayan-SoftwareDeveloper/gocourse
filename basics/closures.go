package main

import "fmt"

func addSequence() func() int {

	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := addSequence()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt2 := addSequence()
	fmt.Println(nextInt2())
	fmt.Println(nextInt())
}
