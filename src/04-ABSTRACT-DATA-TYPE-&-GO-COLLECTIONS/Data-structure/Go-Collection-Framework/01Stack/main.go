/*Stack is a special kind of data structure that follows Last-In-First-Out (LIFO) strategy. This means that the element that
is added last will be the first to be removed.

The various applications of stack are:
1. Recursion: recursive calls are preformed using system stack.
2. Postfix evaluation of expression.
3. Backtracking
4. Depth-first search of trees and graphs.
5. Converting a decimal number into a binary number etc.*/

/*
Stack ADT Operations***********

· Push(k): Adds a new item to the top of the stack
· Pop(): Remove an element from the top of the stack and return its value.
· Top(): Returns the value of the element at the top of the stack
· Size(): Returns the number of elements in the stack
· IsEmpty(): determines whether the stack is empty. It returns 1 if the stack is empty or return 0.
Note: All the above Stack operations are implemented in O(1) Time Complexity.
Stack implementation using Go collections
*/
package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func main() {
	s := stack.New()
	s.Push(2)
	s.Push(3)
	s.Push(4)
	for s.Len() != 0 {
		val := s.Pop()
		fmt.Print(val, " ")
	}

}
