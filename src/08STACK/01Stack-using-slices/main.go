package main

import "fmt"

// Stack using slices
// Implement a stack using slice.
// Example 8.1:
type StackInt struct {
	s []int
}

// If user does not provide the max capacity of the list. Then a list of 1000 elements is created.
// The top is the index to the top of the stack.
// Number of elements in the stack is governed by the “top” index and top is initialized to -1 when a stack is initialized.
// Top index value of -1 indicates that the stack is empty in the beginning.
func (s *StackInt) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

// isEmpty() function returns 1 if stack is empty or 0 in all other cases. By comparing the top index value with -1.

func (s *StackInt) Length() int {
	length := len(s.s)
	return length
}

// size() function returns the number of elements in the stack. It just returns "top+1". As the top is referring the list index
// of the stack top variable so we need to add one to it.

func (s *StackInt) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

// The print function will print the elements of the list.
func (s *StackInt) Push(value int) {
	s.s = append(s.s, value)
}

// push() function checks whether the stack has enough space to store one more element, then it increases the "top" by
// one. Finally sort the data in the stack "data" list. In case, stack is full then "stack overflow" message is printed and that
// value will not be added to the stack and will be ignored.
func (s *StackInt) Pop() int {
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

// In the pop() function, first it will check that there are some elements in the stack by checking its top index. If some
// element is there in the stack, then it will store the top most element value in a variable "value". The top index is reduced
// by one. Finally, that value is returned.

func (s *StackInt) Top() int {
	length := len(s.s)
	res := s.s[length-1]
	return res
}

// top() function returns the value of stored in the top element of stack (does not remove it)

func main() {
	s := new(StackInt)
	length := 10
	for i := 0; i < length; i++ {
		s.Push(i)
	}
	fmt.Println(s.Length())
	for i := 0; i < length; i++ {
		fmt.Print(s.Pop(), " ")
	}
	fmt.Println()
}

// Analysis:
// · The user of the stack will create a stack local variable.
// · Use push() and pop() functions to add / remove variables to the stack.
// · Read the top element using the top() function call.
// · Query regarding size of the stack using size() function call
// · Query if stack is empty using isEmpty() function call

// Stack generic implementation
// Example 8.2:
type Stack struct {
	s []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.s = append(s.s, value)
}

// }
// func (s *Stack) Pop() interface{} {
// 	if s.IsEmpty() {
// 		return nil
// 	}
// 	length := len(s.s)
// 	res := s.s[length-1]
// 	s.s = s.s[:length-1]
// 	return res
// }
func (s *Stack) Top() interface{} {
	length := len(s.s)
	res := s.s[length-1]
	return res
}
