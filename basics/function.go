package main

import "fmt"

func main() {

	//anonymus funciton: function without name

	greet := func() {
		fmt.Println("Hello Anonymus")
	}

	greet()
	q, r := divide(6, 3)
	fmt.Printf("Q:%d, Rem:%d\n", q, r)

	sum(1, 2, 3, 4, 5) //Ellipsis

	numbers := []int{1, 2, 3, 4}
	sum(numbers...)

}

//example of Ellipsis, variadic function, Variadic parameter should be at the end

func sum(nums ...int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println("Sum of nums:=", total)
}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}
