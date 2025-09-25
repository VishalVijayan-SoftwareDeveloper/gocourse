package main

import "fmt"

func main() {
	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println("its an integer")
	case float64:
		fmt.Println("its an float")
	case string:
		fmt.Println("its an string")
	default:
		fmt.Println("its unknown datatype")

	}
}
