package main

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) delete() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}
func main() {
	x, y := 1, 2
	x, y = swap(x, y)
	fmt.Println(x, y)

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(3)
	intStack.push(5)
	fmt.Println(intStack)
	intStack.delete()
	fmt.Println(intStack)

}
