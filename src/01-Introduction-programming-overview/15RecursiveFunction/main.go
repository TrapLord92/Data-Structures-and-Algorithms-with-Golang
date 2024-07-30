package main

import "fmt"

/*A recursive function is a function that calls itself, directly or indirectly.
A recursive method consists of two parts: Termination Condition and Body (which include recursive expansion).
1. Termination Condition: A recursive method always contains one or more terminating condition. A condition in
which recursive method process a simple case and do not call itself.
2. Body (including recursive expansion): The main logic of the recursive method contained in the body of the method.
It also contains the recursion expansion statement that in turn calls the method itself.
Three important properties of recursive algorithm are:
1) A recursive algorithm must have a termination condition.
2) A recursive algorithm must change its state, and move towards the termination condition.
3) A recursive algorithm must call itself.
Note: The speed of a recursive program is slower because of stack overheads. If the same task can be done using an
iterative solution (using loops), then we should prefer an iterative solution in place of recursion to avoid stack
overhead.
Note: Without termination condition, the recursive method may run forever and will finally consume all the stack
memory.*/

//Factorial Calculation. N! = N* (N-1)…. 2*1.
func Factorial(i int) int {
	// Termination Condition
	if i <= 1 {
		return 1
	}
	// Body, Recursive Expansion
	return i * Factorial(i-1)
}

/*Analysis: Factorial() function is calling itself back inside the function. Continue calling itself recursively until you call
the Factorial(1) function.
Time Complexity is O(N)*/
/*******************************************************************************/

/*The Tower of Hanoi (also called the Tower of Brahma) We are given three rods and N number of disks, initially all
the disks are added to first rod (the leftmost one) in decreasing size order. The objective is to transfer the entire stack of
disks from first tower to third tower (the rightmost one), moving only one disk at a time and never a larger one onto a
smaller.*/

func TOHUtil(num int, from string, to string, temp string) {
	if num < 1 {
		return
	}
	TOHUtil(num-1, from, temp, to)
	fmt.Println("Move disk ", num, " from peg ", from, " to peg ", to)
	TOHUtil(num-1, temp, to, from)
}
func TowersOfHanoi(num int) {
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are :")
	TOHUtil(num, "A", "C", "B")
}

/*Analysis: To move N disks from source to destination, we first move N-1 disks from source to temp then move the
lowest Nth disk from source to destination. Then will move N-1 disks from temp to destination.*/

func main() {
	fmt.Println("factorial 5 is :: ", Factorial(5))
	TowersOfHanoi(3)
}
