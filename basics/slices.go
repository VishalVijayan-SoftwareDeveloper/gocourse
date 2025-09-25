package main

import "fmt"

func main() {

	//slice can grow and shrink dynamically

	//var numbers []int
	numbers1 := []int{1, 2, 3, 4, 5}
	//slice1:= make([]int, 5)

	slice1 := numbers1[1:4]
	fmt.Println("Slice1:", slice1)
	slice1 = append(slice1, 6, 7)
	fmt.Println("Slice1:", slice1)

	for index, value := range slice1 {
		fmt.Println(index, value)
	}

}
