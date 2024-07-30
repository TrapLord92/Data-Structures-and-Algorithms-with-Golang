package main

import "fmt"

/*A stack is a memory in which values are stored and retrieved in “last in first out” manner. Data is added to stack using
push operation and data is taken out of stack using pop operation.
1. Initially the stack was empty. Then value 1 is added to stack using push(1) operator.
2. Similarly, push(2) and push(3)
3. Pop operation take the top of the stack. In Stack, data is added and deleted in “last in, first out” manner.
4. First pop() operation will take 3 out of the stack.
5. Similarly, other pop operation will take 2 then 1 out of the stack
6. In the end, the stack is empty when all the elements are taken out of the stack.*/

/*When a method is called, the current execution is stopped and the control goes to the called method. After the called
method exits / returns, the execution resumes from the point at which the execution was stopped.
To get the exact point at which execution should be resumed, the address of the next instruction is stored in the stack.
When the method call completes, the address at the top of the stack is taken out.*/

func function2() {
	fmt.Println("fun2 line 1")
}
func function1() {
	fmt.Println("fun1 line 1")
	function2()
	fmt.Println("fun1 line 2")
}
func main() {
	fmt.Println("main line 1")
	function1()
	fmt.Println("main line 2")
}
