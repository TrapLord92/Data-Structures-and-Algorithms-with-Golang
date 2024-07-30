package main

import "fmt"

type List struct {
	head  *Node
	count int
}
type Node struct {
	value int
	next  *Node
}

// Size of List

func (list *List) Size() int {
	return list.count
}

// IsEmpty function

func (list *List) IsEmpty() bool {
	return (list.count == 0)
}

/*
An element can be inserted into a linked list in various orders. Some of the example cases are mentioned below:
1. Insertion of an element at the start of linked list
2. Insertion of an element at the end of linked list
3. Insertion of an element at the Nth position in linked list
4. Insert element in sorted order in linked list

##Insert element at the Head
Example 7.4:
*/
func (list *List) AddHead(value int) {
	list.head = &Node{value, list.head}
	list.count++
}

/*Analysis:
· We need to create a new node with the value passed to the function as argument.
· While creating the new node the reference stored in head is passed as argument to Node() constructor so that
the next reference will start pointing to the node or null which is referenced by the head node.
· The newly created node will become head of the linked list.
· Size of the list is increased by one.*/

/*###Insertion of an element at the end
Example 7.5: Insertion of an element at the end of linked list*/

func (list *List) addTail(value int) {
	curr := list.head
	newNode := &Node{value, nil}
	if curr == nil {
		list.head = newNode
		return
	}
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

/*Analysis:
· New node is created and the value is stored inside it.
· If the list is empty. Next of new node is null. And head will store the reference to the newly created node.
· If list is not empty then we will traverse until the end of the list.
· Finally, new node is added to the end of the list.*/

/*Note: This operation is un-efficient as each time you want to insert an element you have to traverse to the end of the
list. Therefore, the complexity of creation of the list is n2. So how to make it efficient we have to keep track of the last
element by keeping a tail reference. Therefore, if it is required to insert element at the end of linked list, then we will
keep track of the tail reference also.*/

/*
Traversing Linked List-- means looping
Example 7.6: Print various elements of a linked list
*/
func (list *List) Print() {
	temp := list.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println("")
}

func main() {
	//lst := new(List)
	lst := &List{}
	lst.AddHead(1)
	lst.AddHead(2)
	lst.AddHead(3)
	lst.Print()
}

/*Analysis:
New instance of linked list is created. Various elements are added to list by calling addHead() method.
Finally all the content of list is printed to screen by calling print() method.*/

/*##Sorted Insert
Insert an element in sorted order in linked list given Head reference*/

func (list *List) SortedInsert(value int) {
	newNode := &Node{value, nil}
	curr := list.head
	if curr == nil || curr.value > value {
		newNode.next = list.head
		list.head = newNode
		return
	}
	for curr.next != nil && curr.next.value < value {
		curr = curr.next
	}
	newNode.next = curr.next
	curr.next = newNode
}

/*
Analysis:
· Head of the list is stored in curr.
· A new empty node of the linked list is created. And initialized by storing an argument value into its value. Next of
the node will point to null.
· It checks if the list was empty or if the value stored in the first node is greater than the current value. Then this new
created node will be added to the start of the list. And head need to be modified.
· We iterate through the list to find the proper position to insert the node
· Finally, the node will be added to the list..
·*/
