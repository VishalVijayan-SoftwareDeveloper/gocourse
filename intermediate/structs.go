package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p := Person{"Vishal", "Vijayan", 32}
	fmt.Println(p)
	FullName := p.FullName()
	fmt.Println(FullName)
	fmt.Println(p.Age) //32
	p.IncrementAgeByOne()
	fmt.Println(p.Age) //33
}

func (p Person) FullName() string {
	return p.FirstName + p.LastName
}

func (p *Person) IncrementAgeByOne() {
	p.Age++
}
