// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing container/ring and fmt package
import (
	"container/ring"
	"fmt"
)

/*A circular linked list is a data structure in which the last node is followed by the first node.
The container/ring structures are used to model circular linked lists. An example
implementation of a circular linked list is shown as follows:*/

// main method
func main() {

	var integers []int
	integers = []int{1, 3, 5, 7}

	var circular_list *ring.Ring
	circular_list = ring.New(len(integers))

	var i int
	for i = 0; i < circular_list.Len(); i++ {
		circular_list.Value = integers[i]
		circular_list = circular_list.Next()
	}
	/*The ring.New method with the len n as a parameter creates a circular list of length n. The
	circular linked list is initialized with an integer array by moving through circular_list
	with the Next method. The Do method of ring.Ring class takes the element as an
	interface, and the element is printed as follows:*/

	circular_list.Do(func(element interface{}) {
		fmt.Print(element, ",")
	})
	fmt.Println()
	/*The reverse of the circular list is traversed using the Prev method, and the value is printed
	  in the following code:*/
	for i = 0; i < circular_list.Len(); i++ {
		fmt.Print(circular_list.Value, ",")
		circular_list = circular_list.Prev()
	}
	fmt.Println()
	/*In the following code snippet, the circular list is moved two elements forward using the
	Move method, and the value is printed:*/
	circular_list = circular_list.Move(2)
	circular_list.Do(func(element interface{}) {
		fmt.Print(element, ",")
	})
	fmt.Println()
}
