// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

/*LinkedList
LinkedList is a sequence of nodes that have properties and a reference to the next node in
the sequence. It is a linear data structure that is used to store data. The data structure
permits the addition and deletion of components from any node next to another node. They
are not stored contiguously in memory, which makes them different arrays.
The following sections will look at the structures and methods in a linked list.*/

// Node class
type Node struct {
	property int   //variable
	nextNode *Node //node pointer
}

// LinkedList class
type LinkedList struct {
	headNode *Node
}

// AddToHead method of LinkedList class
/*The AddToHead method adds the node to the start of the linked list. The
AddToHead method of the LinkedList class has a parameter integer property. The
property is used to initialize the node. A new node is instantiated and its property is set to
the property parameter that's passed. The nextNode points to the current headNode of
linkedList, and headNode is set to the pointer of the new node that's created, as shown*/
func (linkedList *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	if linkedList.headNode != nil {
		//fmt.Println(node.property)
		node.nextNode = linkedList.headNode
	}

	linkedList.headNode = node

}

/*In the following code snippet, the NodeWithValue method of LinkedList returns the
node with the property value. The list is traversed and checked to see whether the
property value is equal to the parameter property:*/
// NodeWithValue method returns Node given parameter property
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

/*The AddAfter method adds the node after a specific node. The AddAfter method of
LinkedList has nodeProperty and property parameters. A node with
the nodeProperty value is retrieved using the NodeWithValue method. A node with
property is created and added after the NodeWith node, as follows:*/
// AddAfter method  adds a node with nodeProperty after node with property

func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var nodeWith *Node

	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		//fmt.Println(node.property)
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}

}

// LastNode method returns the last Node
/*The LastNode method of LinkedList returns the node at the end of the list. The list is
traversed to check whether nextNode is nil from nextNode of headNode, as follows:*/
func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd method adds the node with property to the end
/*The AddToEnd method adds the node at the end of the list. In the following code, the
current lastNode is found and its nextNode property is set as the added node:*/
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var lastNode *Node

	lastNode = linkedList.LastNode()

	if lastNode != nil {
		lastNode.nextNode = node
	}

}

// IterateList method iterates over LinkedList
/*The IterateList method of the LinkedList class iterates from the headNode property
and prints the property of the current head node. The iteration happens with the head node
moves to nextNode of the headNode property until the current node is no longer equal to
nil. The following code shows the IterateList method of the LinkedList class:*/
func (linkedList *LinkedList) IterateList() {

	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)

	}
}

// main method
func main() {

	var linkedList LinkedList

	linkedList = LinkedList{}

	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	//fmt.Println(linkedList.headNode.property)

	linkedList.IterateList()

}
