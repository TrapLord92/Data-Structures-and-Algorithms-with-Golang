// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

/*Singly linked lists
A singly linked list is a dynamic data structure in which addition and removal operations
are easy; this is because it's a dynamic data structure and not fixed. Stack and queue data
structures are implemented with linked lists. More memory is consumed when elements
are dynamically added because dynamic data structures aren't fixed. Random retrieval is
not possible with a singly linked list because you will need to traverse the nodes for a
positioned node. Insertion into a singly linked list can be at the beginning or end of the list,
and after a specified node. Deletion can happen at the beginning or end of the list and after
a specified node.
Reversing a singly linked list is shown in this section.*/

/*The Node class is defined in this snippet with a node pointer, nextNode, and a rune
property:*/
// Node struct
type Node struct {
	nextNode *Node
	property rune
}

// create List method
/*The CreateLinkedList method
The CreateLinkedList method creates a linked list of runes from a to z:*/
func CreateLinkedList() *Node {

	var headNode *Node
	headNode = &Node{nil, 'a'}

	var currNode *Node
	currNode = headNode

	var i rune
	for i = 'b'; i <= 'z'; i++ {
		var node *Node
		node = &Node{nil, i}
		currNode.nextNode = node
		currNode = node
	}

	return headNode
}

// Stringify  List method
func StringifyList(nodeList *Node) {

	var currNode *Node
	currNode = nodeList
	for {
		fmt.Printf("%c", currNode.property)

		if currNode.nextNode != nil {
			currNode = currNode.nextNode
		} else {
			break
		}
	}
	fmt.Println("")
}

// Reverse List method
/*The ReverseLinkedList method
The ReverseLinkedList function takes a node pointer, nodeList, and returns a node
pointer to a reversed linked list.
The following code snippet shows how the linked list is reversed:*/
func ReverseLinkedList(nodeList *Node) *Node {

	var currNode *Node
	currNode = nodeList
	var topNode *Node = nil
	for {
		if currNode == nil {
			break
		}
		var tempNode *Node
		tempNode = currNode.nextNode
		currNode.nextNode = topNode
		topNode = currNode
		currNode = tempNode
	}

	return topNode
}

// main method
/*The main method
The main method creates the linked list, and prints the linked list and the reversed list in
string format:*/
func main() {

	var linkedList = CreateLinkedList()
	StringifyList(linkedList)
	StringifyList(ReverseLinkedList(linkedList))
}
