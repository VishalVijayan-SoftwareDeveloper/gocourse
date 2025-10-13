package main

import "fmt"

type Element[T any] struct {
	data T
	next *Element[T]
}

type List[T any] struct {
	head *Element[T]
}

func (l *List[T]) Add(data T) {
	newNode := &Element[T]{data: data}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *List[T]) Display() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ->", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	lst := List[int]{}
	lst.Add(10)
	lst.Add(20)
	lst.Add(30)
	lst.Add(40)
	lst.Add(50)

	lst.Display()

}
