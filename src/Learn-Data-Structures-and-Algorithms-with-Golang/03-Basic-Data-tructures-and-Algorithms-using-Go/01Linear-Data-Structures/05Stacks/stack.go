// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
	"strconv"
)

// Element class
type Element struct {
	elementValue int
}

// String method on Element class
func (element *Element) String() string {
	//fmt.Println(element.elementValue)

	return strconv.Itoa(element.elementValue)
}

/*A stack is a last in, first out structure in which items are added from the top. Stacks are used
in parsers for solving maze algorithms. Push, pop, top, and get size are the typical
operations that are allowed on stack data structures. Syntax parsing, backtracking, and
compiling time memory management are some real-life scenarios where stacks can be used.
An example of stack implementation is as follows */

// NewStack returns a new stack.
/*The New method on the Stack class creates a dynamic array of elements. The Stack class
has the count and array pointer of elements. The code snippet with the Stack class
definition and the New method is as follows:*/
func (stack *Stack) New() {
	stack.elements = make([]*Element, 0)
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	elements     []*Element
	elementCount int
}

// Push adds a node to the stack.
/*The Push method adds the node to the top of the stack class. In the following code
sample, the Push method on the Stack class adds the element to the elements array and
increases the Count element, while the append method adds the element to the elements of
the stack class:*/
func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements[:stack.elementCount], element)
	stack.elementCount = stack.elementCount + 1
}

// Pop removes and returns a node from the stack in last to first order.
/*The Pop method on the Stack implementation removes the last element from the element
array and returns the element, as shown in the following code. The len method returns the
length of the elements array:*/
func (stack *Stack) Pop() *Element {
	if stack.elementCount == 0 {
		return nil
	}

	var length int = len(stack.elements)
	var element *Element = stack.elements[length-1]
	//stack.elementCount = stack.elementCount - 1
	if length > 1 {
		stack.elements = stack.elements[:length-1]

	} else {
		stack.elements = stack.elements[0:]

	}
	stack.elementCount = len(stack.elements)
	return element
}

/*In the following code section, the main method creates a stack, calls the New method, and
pushes the elements after initializing them. The popped-out element value and the order is
printed:*/

// main method
func main() {
	var stack *Stack = &Stack{}
	stack.New()
	var element1 *Element = &Element{3}
	var element2 *Element = &Element{5}
	var element3 *Element = &Element{7}
	var element4 *Element = &Element{9}
	stack.Push(element1)
	stack.Push(element2)
	stack.Push(element3)
	stack.Push(element4)
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
}
