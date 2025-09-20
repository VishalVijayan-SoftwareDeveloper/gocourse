package main

import (
	"fmt"
)

func main() {

	//iterate over a range
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	//iterate over a collection
	number := []int{1, 2, 3, 4, 5}
	for index, value := range number {
		fmt.Printf("Index:%v, Value:%v\n", index, value)
	}

	//display odd number between 1 to 10

	for i := 0; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Printf("Odd number:%v\n", i)
		}

	}

	//Diplay star on Pyramids with 5 rows

	row := 5
	for i := 0; i <= row; i++ {

		for j := 1; j <= row-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// while in go
	var guess int
	var target int = 77
	for {
		fmt.Println("Enter your guess:")
		fmt.Scanln(&guess)
		if guess == target {
			fmt.Println(" Congratulations!!!  Yoou guessed correct number")
			break
		} else if guess < target {
			fmt.Println("Too low and try again")
		} else {
			fmt.Println("Too high and try again")
		}
	}

}
