package main

import "fmt"

// Problems in Stack
// Balanced Parenthesis
// Example 8.4: Stacks can be used to check a program for balanced symbols (such as {}, (), []). The closing symbol
// should be matched with the most recently seen opening symbol.
// Example: {()} is legal, {() ({})} is legal, but {((} and {(}) are not legal

type Stack struct {
	s []interface{}
}
type StackInt struct {
	s []int
}

func (s *Stack) Push(value interface{}) {
	s.s = append(s.s, value)
}
func (s *Stack) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}
func (s *Stack) Top() interface{} {
	length := len(s.s)
	res := s.s[length-1]
	return res
}

func IsBalancedParenthesis(expn string) bool {
	stk := new(Stack)
	for _, ch := range expn {
		switch ch {
		case '{', '[', '(':
			stk.Push(ch)
		case '}':
			val := stk.Pop()
			if val != '{' {
				return false
			}
		case ']':
			val := stk.Pop()
			if val != '[' {
				return false
			}
		case ')':
			val := stk.Pop()
			if val != '(' {
				return false
			}
		}
	}
	return stk.IsEmpty()
}
func main() {
	expn := "{()}[]"
	value := IsBalancedParenthesis(expn)
	fmt.Println("Given Expn:", expn)
	fmt.Println("Result aft1er isParenthesisMatched:", value)
}

// Analysis:
// · Traverse the input string when we get an opening parenthesis we push it into stack. Moreover, when we get a
// closing parenthesis then we pop a parenthesis from the stack and compare if it is the corresponding to the one on
// the closing parenthesis.
// · We return false if there is a mismatch of parenthesis.
// · If at the end of the whole staring traversal, we reached to the end of the string and the stack is empty then we have
// balanced parenthesis.
