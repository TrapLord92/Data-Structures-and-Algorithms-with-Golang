// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and bytes package
import (
	"fmt"
)

/*Recursion is an algorithm in which one of the steps invokes the currently running method
or function. This algorithm acquires the outcome for the input by applying basic tasks and
then returns the value.During recursion, if the base
condition is not reached, then a stack overflow condition may arise.*/

/*A recursion algorithm is implemented in the following code snippet. The Factor method
takes the num as a parameter and returns the factorial of num. The method uses recursion to
calculate the factorial of the number:*/
//factorial method
func Factor(num int) int {
	if num <= 1 {
		return 1
	}
	return num * Factor(num-1)
}

// main method
func main() {
	var num int = 5
	fmt.Println("Factorial: %d is %d", num, Factor(num))
}
