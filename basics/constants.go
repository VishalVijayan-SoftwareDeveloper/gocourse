package main

import "fmt"

const PI = 3.14
const GRAVITY = 9.8

// Pascal case naming
type Empolyee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var user_id int = 1 // snake case convention
	fmt.Println(user_id)
	const (
		monday  = 1
		tuesday = 2
	)

}
