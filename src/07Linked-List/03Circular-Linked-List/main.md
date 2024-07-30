package main

import "fmt"

/*Circular Linked List
This type is similar to the singly linked list except that the last element points to the first node of the list. The link
portion of the last node contains the address of the first node.
Example 7.42:*/
type CircularLinkedList struct {
	tail  *Node
	count int
}
type Node struct {
	value int
	next  *Node
}

func (list *CircularLinkedList) Size() int {
	return list.count
}
func (list *CircularLinkedList) IsEmpty() bool {
	return list.count == 0
}
func (list *CircularLinkedList) Peek() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListError")
		return 0, false
	}
	return list.tail.next.value, true
}

/*Analysis: In the circular linked list, we just need the pointer to the tail node. As head node can be easily reached from
tail node. Size(), isEmpty() and peek() functions remains the same.
Insert element in front
Example 7.43:*/
func (list *CircularLinkedList) AddHead(value int) {
	temp := &Node{value, nil}
	if list.IsEmpty() {
		list.tail = temp
		temp.next = temp
	} else {
		temp.next = list.tail.next
		list.tail.next = temp
	}
	list.count++
}

/*Analysis:
· First, we create node with given value and its next pointing to null.
· If the list is empty then tail of the list will point to it. In addition, the next of node will point to itself
· If the list is not empty then the next of the new node will be next of the tail. In addition, tail next will start pointing
to the new node.
· Thus, the new node is added to the head of the list.
· The demo program create an instance of CircularLinkedList class. Then add some value to it and finally print the
content of the list.
Insert element at the end
Example 7.44:*/
func (list *CircularLinkedList) AddTail(value int) {
	temp := &Node{value, nil}
	if list.IsEmpty() {
		list.tail = temp
		temp.next = temp
	} else {
		temp.next = list.tail.next
		list.tail.next = temp
		list.tail = temp
	}
	list.count++
}

/*Analysis: Adding node at the end is same as adding at the beginning. Just need to modify tail reference in place of the
head reference.
Remove element in the front
Example 7.45:*/
func (list *CircularLinkedList) RemoveHead() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListError")
		return 0, false
	}
	value := list.tail.next.value
	if list.tail == list.tail.next {
		list.tail = nil
	} else {
		list.tail.next = list.tail.next.next
	}
	list.count--
	return value, true
}

/*Analysis:
· If the list is empty, then false will be returned.
· The value stored in head is stored in local variable value.
· If tail is equal to its next node that means there is only one node in the list so the tail will become null.
· In all the other cases, the next of tail will point to next element of the head.
· Finally, the value is returned.
Search element in the list
Example 7.46:*/
func (list *CircularLinkedList) IsPresent(data int) bool {
	temp := list.tail
	for i := 0; i < list.count; i++ {
		if temp.value == data {
			return true
		}
		temp = temp.next
	}
	return false
}

/*Analysis: Iterate through the list to find if particular value is there or not.
Print the content of list
Example 7.47:*/
func (list *CircularLinkedList) Print() {
	if list.IsEmpty() {
		return
	}
	temp := list.tail.next
	for temp != list.tail {
		fmt.Println(temp.value, " ")
		temp = temp.next
	}
	fmt.Println(temp.value)
}

/*Analysis: In circular list, tail is used to check end of the list.
Delete List
Example 7.48:*/
func (list *CircularLinkedList) FreeList() {
	list.tail = nil
	list.count = 0
}

/*Analysis: The reference to the list is tail. By making tail null, the whole list is deleted.
Delete a node given its value
Example 7.49:*/
func (list *CircularLinkedList) RemoveNode(key int) bool {
	if list.IsEmpty() {
		return false
	}
	prev := list.tail
	curr := list.tail.next
	head := list.tail.next
	if curr.value == key { // head and single node case.
		if curr == curr.next { // single node case
			list.tail = nil
		} else { // head case
			list.tail.next = list.tail.next.next
		}
		return true
	}
	prev = curr
	curr = curr.next
	for curr != head {
		if curr.value == key {
			if curr == list.tail {
				list.tail = prev
			}
			prev.next = curr.next
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}

/*Analysis: Find the node that need to free. Only difference is that while traversing the list end of list is tracked by the
head reference in place of null.
Copy List Reversed
Example 7.50:*/
func (list *CircularLinkedList) CopyListReversed() *CircularLinkedList {
	cl := new(CircularLinkedList)
	curr := list.tail.next
	head := curr
	if curr != nil {
		cl.AddHead(curr.value)
		curr = curr.next
	}
	for curr != head {
		cl.AddHead(curr.value)
		curr = curr.next
	}
	return cl
}

/*Analysis: The list is traversed and nodes are added to new list at the beginning. There by making the new list reverse of
the given list.
Copy List
Example 7.51:*/
func (list *CircularLinkedList) CopyList() *CircularLinkedList {
	cl := new(CircularLinkedList)
	curr := list.tail.next
	head := curr
	if curr != nil {
		cl.addTail(curr.value)
		curr = curr.next
	}
	for curr != head {
		cl.addTail(curr.value)
		curr = curr.next
	}
	return cl
}

/*Analysis: List is traversed and nodes are added to the new list at the end. There by making the list whose value are
same as the input list.
Doubly Circular list
1. For any linked list there are only three cases zero element, one element, general case
2. To doubly linked list we have a few more things
a) null values
b) Only element (it generally introduces an if statement with null)
c) Always an “if” before “while”. Which will check from this head.
d) General case (check with the initial head kept)
e) Avoid using recursion solutions it makes life harder
Example 7.52:*/
type DoublyCircularLinkedList struct {
	head  *Node
	tail  *Node
	count int
}
type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (
	list *DoublyCircularLinkedList) Size() int {
	return list.count
}
func (
	list *DoublyCircularLinkedList) IsEmpty() bool {
	return list.count == 0
}
func (
	list *DoublyCircularLinkedList) peekHead() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListError")
		return 0, false
	}
	return list.head.value, true
}

// Search value
// Example 7.53:
func (list *DoublyCircularLinkedList) IsPresent(key int) bool {
	temp := list.head
	if list.head == nil {
		return false
	}
	for true {
		if temp.value == key {
			return true
		}
		temp = temp.next
		if temp == list.head {
			break
		}
	}
	return false
}

/*Analysis: Traverse through the list and see if given key is present or not. We if statement to terminate the loop.
Delete list
Example 7.54:*/
func (list *DoublyCircularLinkedList) FreeList() {
	list.head = nil
	list.tail = nil
	list.count = 0
}

/*Analysis: Remove the reference and list will be freed.
Print List
Example 7.55:*/
func (list *DoublyCircularLinkedList) Print() {
	if list.IsEmpty() {
		return
	}
	fmt.Println("List size is ::", list.count)
	fmt.Print("List content :: ")
	temp := list.head
	for true {
		fmt.Print(temp.value, " ")
		temp = temp.next
		if temp == list.head {
			break
		}
	}
	fmt.Println()
}

/*Analysis: Traverse the list and print its content. For loop with if condition to exit loop is used as we want to terminate
when temp is head. Moreover, want to process head node once.
Insert Node at head
Example 7.56: Insert value at the front of the list.*/
func (list *DoublyCircularLinkedList) AddHead(value int) {
	newNode := new(Node)
	newNode.value = value
	if list.count == 0 {
		list.tail = newNode
		list.head = newNode
		newNode.next = newNode
		newNode.prev = newNode
	} else {
		newNode.next = list.head
		newNode.prev = list.head.prev
		list.head.prev = newNode
		newNode.prev.next = newNode
		list.head = newNode
	}
	list.count++
}

/*Analysis:
· A new node is created and if the list is empty then head and tail will point to it. The newly created newNode’s next
and prev also point to newNode.
· If the list is not empty then the pointers are adjusted and a new node is added to the front of the list. Only head need
to be changed in this case.
· Size of the list is increased by one.
Insert Node at tail
Example 7.57:*/
func (list *DoublyCircularLinkedList) AddTail(value int) {
	newNode := new(Node)
	newNode.value = value
	if list.count == 0 {
		list.head = newNode
		list.tail = newNode
		newNode.next = newNode
		newNode.prev = newNode
	} else {
		newNode.next = list.tail.next
		newNode.prev = list.tail
		list.tail.next = newNode
		newNode.next.prev = newNode
		list.tail = newNode
	}
	list.count++
}

/*Analysis:
· A new node is created and if the list is empty then head and tail will point to it. The newly created newNode’s next
and prev also point to newNode.
· If the list is not empty then the pointers are adjusted and a new node is added to the end of the list. Only tail need to
be changed in this case.
· Size of the list is increased by one.
Delete head node
Example 7.58:*/
func (list *DoublyCircularLinkedList) RemoveHead() (int, bool) {
	if list.count == 0 {
		fmt.Println("EmptyListError")
		return 0, false
	}
	value := list.head.value
	list.count--
	if list.count == 0 {
		list.head = nil
		list.tail = nil
		return value, true
	}
	next := list.head.next
	next.prev = list.tail
	list.tail.next = next
	list.head = next
	return value, true
}

/*Analysis: Delete node in a doubly circular linked list is just same as delete node in a circular linked list. Just few extra
next reference need to be adjusted.
Delete tail node
Example 7.59:*/
func (list *DoublyCircularLinkedList) RemoveTail() (int, bool) {
	if list.count == 0 {
		fmt.Println("EmptyListError")
		return 0, false
	}
	value := list.tail.value
	list.count--
	if list.count == 0 {
		list.head = nil
		list.tail = nil
		return value, true
	}
	prev := list.tail.prev
	prev.next = list.head
	list.head.prev = prev
	list.tail = prev
	return value, true
}

/*Analysis: Delete node in a doubly circular linked list is just same as delete node in a circular linked list. Just few extra
prev reference need to be adjusted.*/
