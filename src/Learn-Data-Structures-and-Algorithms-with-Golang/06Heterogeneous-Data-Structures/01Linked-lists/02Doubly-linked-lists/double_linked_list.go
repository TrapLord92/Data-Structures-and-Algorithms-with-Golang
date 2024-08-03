// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and list package
import (
	"container/list"
	"fmt"
)

/*Doubly linked lists
A doubly linked list is a data structure that consists of nodes that have links to the previous
and the next nodes. Doubly linked lists were presented with code examples in Chapter 3,
Linear Data Structures. Lists in Go are implemented as doubly linked lists. The elements 14
and 1 are pushed backward and forward, respectively. The elements 6 and 5 are inserted
before and after. The doubly linked list is iterated and the elements are printed. The code in
this section shows how lists can be used:*/
// main method
func main() {
	var linkedList *list.List
	linkedList = list.New()
	var element *list.Element
	element = linkedList.PushBack(14)

	var frontElement *list.Element
	frontElement = linkedList.PushFront(1)
	linkedList.InsertBefore(6, element)
	linkedList.InsertAfter(5, frontElement)

	var currElement *list.Element
	for currElement = linkedList.Front(); currElement != nil; currElement = currElement.Next() {
		fmt.Println(currElement.Value)
	}

}
