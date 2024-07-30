package main

import "fmt"

// Stack using linked list
// Example 8.3: Implement stack using a linked list.
type Node struct {
	value int
	next  *Node
}
type Stack struct {
	head *Node
	size int
}

func (s *Stack) Size() int {
	return s.size
}
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyError")
		return 0, false
	}
	return s.head.value, true
}
func (s *Stack) Push(value int) {
	s.head = &Node{value, s.head}
	s.size++
}
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyError")
		return 0, false
	}
	value := s.head.value
	s.head = s.head.next
	s.size--
	return value, true
}
func (s *Stack) Print() {
	temp := s.head
	fmt.Print("Value stored in stck are :: ")
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

// Analysis:
// · Stack implemented using a linked list is simply insertion and deletion at the head of a singly linked list.
// · In push() function, memory is created for one node. Then the value is stored into that node. Finally, the node is
// inserted at the beginning of the list.
// · In pop() function, the head of the linked list starts pointing to the second node there by releasing the memory
// allocated to the first node (Garbage collection).
