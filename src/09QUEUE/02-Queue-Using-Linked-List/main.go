package main

import "fmt"

// Queue Using linked list
// Example 9.2:
type Node struct {
	value int
	next  *Node
}
type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) Size() int {
	return q.size
}
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
func (q *Queue) Peek() (int, bool) {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyError")
		return 0, false
	}
	return q.head.value, true
}
func (q *Queue) Print() {
	temp := q.head
	for temp != nil {
		fmt.Println(temp.value, " ")
		temp = temp.next
	}
}

/*Add
Enqueue into a queue using linked list. Nodes are added to the end of the linked list. Below diagram indicates how a
new node is added to the list. The tail is modified every time when a new value is added to the queue. However, the
head is also updated in the case when there is no element in the queue and when that first element is added to the queue
both head and tail will be pointing to it.*/

// func (q *Queue) Add(value int) {
// 	temp := &Node{value}
// 	if q.head == nil {
// 		q.head = temp
// 		q.tail = temp
// 	} else {
// 		q.tail.next = temp
// 		q.tail = temp
// 	}
// 	q.size++
// }

// Analysis: add operation add one element at the end of the Queue (linked list).

/*Remove
In this we need the tail reference as it may be the case there was only one element in the list and the tail reference will
also be modified in case of the remove.
Example 9.4:*/

func (q *Queue) Remove() (int, bool) {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyError")
		return 0, false
	}
	value := q.head.value
	q.head = q.head.next
	q.size--
	return value, true
}

/*Analysis: Remove operation removes first node from the start of the queue (linked list).*/
