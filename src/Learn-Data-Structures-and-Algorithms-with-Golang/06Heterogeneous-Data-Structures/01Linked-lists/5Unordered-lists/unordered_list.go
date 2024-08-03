package main

//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book

/*Unordered lists
An unordered list is implemented as a linked list. In an unordered list, the relative
positions of items in contiguous memory don't need to be maintained. The values will be
placed in a random fashion.
An unordered list starts with a <ul> tag in HTML 5.0. Each list item is coded with <li>
tags. Here's an example:
<ul>
<li> First book </li>
<li> Second book </li>
<li> Third book </li>
</ul>
The following is an example of an unordered list in Golang. The Node class has a property
and a nextNode pointer, as shown in the following code. The linked list will have a set of
nodes with a property attribute. The unordered list is presented in the script called
unordered_list.go:
//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main
// importing fmt package
import (
"fmt"
)*/
//Node class
import (
	"fmt"
)

// Node class
type Node struct {
	property int
	nextNode *Node
}

// UnOrderedList class
/*The UnOrderedList class
The unordered list consists of elements that are not ordered by numbers. An
UnOrderedList class has a headNode pointer as the property. Traversing to the next node
from the head node, you can iterate through the linked list:*/
type UnOrderedList struct {
	headNode *Node
}

// AddToHead method of UnOrderedList class
/*The AddtoHead method
The AddtoHead method adds the node to the head of the unordered list. The AddToHead
method of the UnOrderedList class has a property parameter that's an integer. It will
make the headNode point to a new node created with property, and the nextNode points
to the current headNode of the unordered list:*/
func (UnOrderedList *UnOrderedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	if UnOrderedList.headNode != nil {
		node.nextNode = UnOrderedList.headNode
	}

	UnOrderedList.headNode = node

}

// IterateList method iterates over UnOrderedList
/*The IterateList method
The IterateList method of the UnOrderedList class prints the node property of the
nodes in the list. This is shown in the following code:*/
func (UnOrderedList *UnOrderedList) IterateList() {

	var node *Node
	for node = UnOrderedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)

	}
}

// main method
/*The main method
The main method creates an instance of a linked list, and integer properties 1, 3, 5, and 7
are added to the head of the linked list. The linked list's headNode property is printed after
the elements are added:*/
func main() {

	var unOrderedList UnOrderedList

	unOrderedList = UnOrderedList{}

	unOrderedList.AddToHead(1)
	unOrderedList.AddToHead(3)
	unOrderedList.AddToHead(5)
	unOrderedList.AddToHead(7)
	unOrderedList.IterateList()

}
