package main

import "fmt"

func main() {
	process()
}

func process() {
	defer fmt.Println("Deffered statement executed")
	defer fmt.Println("second deferred statement")
	fmt.Println("HEllo world")
}
