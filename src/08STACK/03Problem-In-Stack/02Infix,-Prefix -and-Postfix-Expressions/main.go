package main

import (
	"fmt"
	"strconv"
	"strings"

)

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

func InfixToPostfix(expn string) string {
	fmt.Println(expn)
	stk := new(Stack)
	output := ""
	for _, ch := range expn {
		if ch <= '9' && ch >= '0' {
			output = output + string(ch)
		} else {
			switch ch {
			case '+', '-', '*', '/', '%', '^':
				for stk.IsEmpty() == false && precedence(ch) <= precedence(stk.Top().(rune)) {
					out := stk.Pop().(rune)
					output = output + " " + string(out)
				}
				stk.Push(ch)
				output = output + " "
			case '(':
				stk.Push(ch)
			case ')':
				for stk.IsEmpty() == false && stk.Top() != '(' {
					out := stk.Pop().(rune)
					output = output + " " + string(out) + " "
				}
				if stk.IsEmpty() == false && stk.Top() == '(' {
					stk.Pop()
				}
			}
		}
	}
	for stk.IsEmpty() == false {
		out := stk.Pop().(rune)
		output = output + string(out) + " "
	}
	return output
}
func precedence(x rune) int {
	if x == '(' {
		return (0)
	}
	if x == '+' || x == '-' {
		return (1)
	}
	if x == '*' || x == '/' || x == '%' {
		return (2)
	}
	if x == '^' {
		return (3)
	}
	return (4)
}

func main() {
	expn := "10+((3))*5/(16-4)"
	value := InfixToPostfix(expn)
	fmt.Println("Infix Expn: ", expn)
	fmt.Println("Postfix Expn: ", value)
	main1()
	main2()
}

// Analysis:
// · Print operands in the same order as they arrive.
// · If the stack is empty or contains a left parenthesis “(” on top, we should push the incoming operator in the stack.
// · If the incoming symbol is a left parenthesis ”(”, push left parenthesis in the stack.
// · If the incoming symbol is a right parenthesis “)”, pop from the stack and print the operators until you see a left
// parenthesis “)”. Discard the pair of parentheses.
// · If the precedence of incoming symbol is higher than the precedence of operator at the top of the stack, then push it
// to the stack.

// · If the incoming symbol has, an equal precedence compared to the top of the stack, use association. If the
// association is left to right, then pop and print the symbol at the top of the stack and then push the incoming
// operator. If the association is right to left, then push the incoming operator.
// · If the precedence of incoming symbol is lower than the precedence of operator on the top of the stack, then pop and
// print the top operator. Then compare the incoming operator against the new operator at the top of the stack.
// · At the end of the expression, pop and print all operators on the stack.

func InfixToPrefix(expn string) string {
	expn = reverseString(expn)
	expn = replaceParanthesis(expn)
	expn = InfixToPostfix(expn)
	expn = reverseString(expn)
	return expn
}
func reverseString(in string) string {
	expn := []rune(in)
	lower := 0
	upper := len(expn) - 1
	for lower < upper {
		expn[lower], expn[upper] = expn[upper], expn[lower]
		lower++
		upper--
	}
	return string(expn)
}
func replaceParanthesis(str string) string {
	a := []rune(str)
	lower := 0
	upper := len(a) - 1
	for lower <= upper {
		if a[lower] == '(' {
			a[lower] = ')'
		} else if a[lower] == ')' {
			a[lower] = '('
		}
		lower++
	}
	return string(a)
}
func main1() {
	expn := "10+((3))*5/(16-4)"
	value := InfixToPrefix(expn)
	fmt.Println("Infix Expn: ", expn)
	fmt.Println("Prefix Expn: ", value)
}

// Analysis:
// 1. Reverse the given infix expression.
// 2. Replace '(' with ')' and ')' with '(' in the reversed expression.
// 3. Now, apply infix to postfix subroutine already discussed.
// 4. Reverse the generated postfix expression and this will give required prefix expression.

// Postfix Evaluate
// Write a postfixEvaluate() function to evaluate a postfix expression. Such as: 1 2 + 3 4 + *
// Example 8.7:
func postfixEvaluate(expn string) int {
	stk := new(Stack)
	str := strings.Split(expn, " ")
	for _, tkn := range str {
		value, err := strconv.Atoi(tkn)
		if err == nil {
			stk.Push(value)
		} else {
			num1 := stk.Pop().(int)
			num2 := stk.Pop().(int)
			switch tkn {
			case "+":
				stk.Push(num1 + num2)
			case "-":
				stk.Push(num1 - num2)
			case "*":
				stk.Push(num1 * num2)
			case "/":
				stk.Push(num1 / num2)
			}
		}
	}
	return stk.Pop().(int)
}
func main2() {
	expn := "6 5 2 3 + 8 * + 3 + *"
	value := postfixEvaluate(expn)
	fmt.Println("Given Postfix Expn: ", expn)
	fmt.Println("Result after Evaluation: ", value)
}

// Analysis:
// 1) Create a stack to store values or operands.
// 2) Scan through the given expression and do following for each element:
// a) If the element is a number, then push it into the stack.
// b) If the element is an operator, then pop values from the stack. Evaluate the operator over the values and push
// the result into the stack.
// 3) When the expression is scanned completely, the number in the stack is the result.
