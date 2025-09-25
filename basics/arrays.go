package main

import "fmt"

func main() {
	var numbers [5]int
	fmt.Println(numbers) //output:[0 0 0 0 0]
	numbers[0] = 1
	numbers[4] = 5
	//sring array
	fruits := []string{"apple", "banane", "orange"}
	fmt.Println(fruits)

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	//compare arrays

	array1 := [4]int{1, 2, 3, 4}
	array2 := [4]int{4, 5, 6, 7}

	fmt.Println("if Array1 is equual to array2:", array1 == array2)
}
