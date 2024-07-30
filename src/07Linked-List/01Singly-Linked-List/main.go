package main

/*Let us look at the Node. The value part of node is of type integer, but it can be some other data-type. The link part of
node is named as next in the below class definition.
Note: For a singly linked, we should always test these three test cases before saying that the code is good to go. This
one node and zero node case is used to catch boundary cases. It is always to take care of these cases before submitting
code to the reviewer.
· Zero element / Empty linked list.
· One element / Just single node case.
· General case.*/

/*The various basic operations that we can perform on linked lists, many of these operations require list traversal:
· Insert an element in the list, this operation is used to create a linked list.
· Print various elements of the list.
· Search an element in the list.
· Delete an element from the list.
· Reverse a linked list.
You cannot use Head to traverse a linked list because if we use the head, then we lose the nodes of the list. We have to
use another reference variable of same data-type as the head.*/

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
