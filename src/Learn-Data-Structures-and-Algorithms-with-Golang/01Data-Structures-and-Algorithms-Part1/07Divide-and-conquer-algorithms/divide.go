// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

/*A divide and conquer algorithm breaks a complex problem into smaller problems and
solves these smaller problems. The smaller problem will be further broken down till it is a
known problem. The approach is to recursively solve the sub-problems and merge the
solutions of the sub-problems.
Recursion, quick sort, binary search, fast Fourier transform, and merge sort are good
examples of divide and conquer algorithms. Memory is efficiently used with these
algorithms. Performance is sometimes an issue in the case of recursion. On multiprocessor
machines, these algorithms can be executed on different processors after breaking them
down into sub-problems. A divide and conquer algorithm is shown in the following code:*/
// fibonacci method given k integer
func fibonacci(k int) int {

	if k <= 1 {
		return 1
	}
	return fibonacci(k-1) + fibonacci(k-2)

}

// main method
func main() {

	var m int = 5

	for m = 0; m < 8; m++ {

		var fib = fibonacci(m)
		fmt.Println(fib)
	}

}
