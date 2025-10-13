package main

import "fmt"

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func PrintLinkedList(head *LinkedList) {

	for head != nil {
		fmt.Print(head.Val, "-->")
		head = head.Next
	}
	fmt.Println("nil")
}

func ReverseLinkedList(head *LinkedList) *LinkedList {

	var prev *LinkedList
	current := head

	for current != nil {
		fmt.Printf("current: %d\n", current.Val)

		next := current.Next
		current.Next = prev
		prev = current
		fmt.Printf("previous: %d\n", prev.Val)
		current = next

	}
	PrintLinkedList(prev)
	return prev
}

/*
func Add(data int)  {
	newNode:=&LinkedList{}
	if new
}

*/

func main() {

	head := &LinkedList{Val: 1}
	head.Next = &LinkedList{Val: 2}
	head.Next.Next = &LinkedList{Val: 3}
	head.Next.Next.Next = &LinkedList{Val: 4}

	PrintLinkedList(head)

	ReverseLinkedList(head)

}
