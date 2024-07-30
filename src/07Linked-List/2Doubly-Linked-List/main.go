package main

import "fmt"

/*Doubly Linked List
In a Doubly Linked list, there are two references in each node. These references are called prev and next. The prev
reference of the node will point to the node before it and the next reference will point to the node next to the given
node.
Let us look at the Node. The value part of the node is of type integer, but it can be of some other data-type. The two
link references are prev and next.
Search in a single linked list can be only done in one direction. Since all elements in the list has reference to the next
item in the list. Therefore, traversal of linked list is linear in nature. In a doubly linked list, we keep track of both head
of the linked list and tail of linked list.
In doubly linked list linked list below are few cases that we need to keep in mind while coding:
· Zero element case (head and tail both can be modified)
· Only element case (head and tail both can be modified)
· First element (head can be modified)
· General case
· The last element (tail can be modified)
Note: Any program that is likely to change head reference or tail reference is to be passed as a double reference, which
is pointing to head or tail reference.
Basic operations of Linked List
Basic operation of a linked list requires traversing a linked list. The various operations that we can perform on linked
lists, many of these operations require list traversal:
1. Insert an element in the list, this operation is used to create a linked list.
2. Print various elements of the list.
3. Search an element in the list.
4. Delete an element from the list.
5. Reverse a linked list.
For doubly linked list, we have following cases to consider:
1. null values (head and tail both can be modified)
2. Only element (head and tail both can be modified)
3. First element (head can be modified)
4. General case
5. Last element (tail can be modified)*/
//Creating strutura para double linked list
type DoublyLinkedList struct {
	head  *Node
	tail  *Node
	count int
}
type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (list *DoublyLinkedList) Size() int {
	return list.count
}
func (list *DoublyLinkedList) IsEmpty() bool {
	return list.count == 0
}
func (list *DoublyLinkedList) Peek() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListError")
		return 0, false
	}
	return list.head.value, true
}

// Insert at Head

func (list *DoublyLinkedList) AddHead(value int) {
	newNode := &Node{value, nil, nil}
	if list.count == 0 {
		list.tail = newNode
		list.head = newNode
	} else {
		list.head.prev = newNode
		newNode.next = list.head
		list.head = newNode
	}
	list.count++
}

/*Analysis: Insert in double linked list is same as insert in a singly linked list.
· Create a node assign null to prev reference of the node.
· If the list is empty then tail and head will point to the new node.
· If the list is not empty then prev of head will point to newNode and next of newNode will point to head. Then head
will be modified to point to newNode.*/
// Insert at Tail
// Example 7.31: Insert an element at the end of the list.
func (list *DoublyLinkedList) AddTail(value int) {
	newNode := &Node{value, nil, nil}
	if list.count == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
	list.count++
}

/*Analysis: Find the proper location of the node and add it to the list. Manage next and prev reference of the node so that
list always remain double linked list.
Remove Head of doubly linked list*/
// Example 7.32:
func (list *DoublyLinkedList) RemoveHead() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListError")
		return 0, false
	}
	value := list.head.value
	list.head = list.head.next
	if list.head == nil {
		list.tail = nil
	} else {
		list.head.prev = nil
	}
	list.count--
	return value, true
}

/*Analysis:
· If the list is empty then EmptyListError message will be printed.
· Now head will point to its next.
· If head is null then this was single node list case, tail also need to be made null.
· In all the general case head. Prev will be set to null.
· Size of list will be reduced by one and value of node is returned.*/

// Delete a node given its value
// Example 7.33: Delete node in linked list
func (list *DoublyLinkedList) RemoveNode(key int) bool {
	curr := list.head
	if curr == nil { // empty list
		return false
	}
	if curr.value == key { // head is the node with value key.
		curr = curr.next
		list.count--
		if curr != nil {
			list.head = curr
			list.head.prev = nil
		} else {
			list.tail = nil // only one element in list.
		}
		return true
	}
	for curr.next != nil {
		if curr.next.value == key {
			curr.next = curr.next.next
			if curr.next == nil { // last element case.
				list.tail = curr
			} else {
				curr.next.prev = curr
			}
			list.count--
			return true
		}
		curr = curr.next
	}
	return false
}

/*Analysis: Traverse the list find the node, which need to be deleted. Then remove it and adjust next reference of the
node before it and prev reference of the node next to it.*/
// Search list
// Example 7.34:
func (list *DoublyLinkedList) IsPresent(key int) bool {
	temp := list.head
	for temp != nil {
		if temp.value == key {
			return true
		}
		temp = temp.next
	}
	return false
}

/*Analysis: Traverse the list and find if some value is resent or not.
Free List
Example 7.35:*/
func (list *DoublyLinkedList) FreeList() {
	list.tail = nil
	list.head = nil
	list.count = 0
}

/*Analysis: Just head and tail references need to point to null. The rest of the list will automatically deleted by garbage
collection.

Print list
Example 7.36:*/
func (list *DoublyLinkedList) Print() {
	temp := list.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

/*Analysis: Traverse the list and print the value of each node.
Reverse a doubly linked List iteratively
Example 7.37:*/
func (list *DoublyLinkedList) ReverseList() {
	curr := list.head
	var tempNode *Node
	for curr != nil {
		tempNode = curr.next
		curr.next = curr.prev
		curr.prev = tempNode
		if curr.prev == nil {
			list.tail = list.head
			list.head = curr
			return
		}
		curr = curr.prev
	}
	return
}

/*
Analysis: Traverse the list. Swap the next and prev. then traverse to the direction curr.prev, which was next before
swap. If you reach the end of the list then set head and tail.

Copy List Reversed
Example 7.38: Copy the content of the list into another list in reverse order.*/
func (list *DoublyLinkedList) CopyListReversed(dll *DoublyLinkedList) {
	curr := list.head
	for curr != nil {
		dll.AddHead(curr.value)
		curr = curr.next
	}
}

/*Analysis:
· Create a DoublyLinkedList class object dll.
· Traverse through the list and copy the value of the nodes into another list by calling addHead() method.
· Since the new nodes are added to the head of the list, the new list formed have nodes order reverse there by making
reverse list.

Copy List
Example 7.39:*/
func (list *DoublyLinkedList) CopyList(dll *DoublyLinkedList) {
	curr := list.head
	for curr != nil {
		dll.AddTail(curr.value)
		curr = curr.next
	}
}

/*Analysis:
· Create a DoublyLinkedList class object dll.
· Traverse through the list and copy the value of the nodes into another list by calling addTail() method.
· Since the new nodes are added to the tail of the list, the new list formed have nodes order same as the original list.
Sorted Insert
Example 7.40:*/
func (list *DoublyLinkedList) SortedInsert(value int) {
	temp := &Node{value, nil, nil}
	curr := list.head
	if curr == nil { // first element
		list.head = temp
		list.tail = temp
	}
	if list.head.value <= value { // at the begining
		temp.next = list.head
		list.head.prev = temp
		list.head = temp
	}
	for curr.next != nil && curr.next.value > value { // treversal
		curr = curr.next
	}
	if curr.next == nil { // at the end
		list.tail = temp
		temp.prev = curr
		curr.next = temp
	} else { // all other
		temp.next = curr.next
		temp.prev = curr
		curr.next = temp
		temp.next.prev = temp
	}
}

/*Analysis:
· We need to consider only element case first. In this case, both head and tail will modify.
· Then we need to consider the case when head will be modified when new node is added to the beginning of the list.
· Then we need to consider general cases
· Finally, we need to consider the case when tail will be modified.
Remove Duplicate
Example 7.41: Consider the list as sorted remove the repeated value nodes of the list.
# Remove Duplicate*/
func (list *DoublyLinkedList) RemoveDuplicate() {
	curr := list.head
	var deleteMe *Node
	for curr != nil {
		if (curr.next != nil) && curr.value == curr.next.value {
			deleteMe = curr.next
			curr.next = deleteMe.next
			curr.next.prev = curr
			if deleteMe == list.tail {
				list.tail = curr
			}
		} else {
			curr = curr.next
		}
	}
}

/*Analysis:
· Remove duplicate is same as single linked list case.
· Head can never modify only the tail can modify when the last node is*/
