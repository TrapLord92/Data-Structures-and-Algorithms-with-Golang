// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

/*Spiral matrix
A spiral matrix is an arrangement of n x n integers in which integers are arranged spirally
in sequentially increasing order. A spiral matrix is an old toy algorithm. The spiral order is
maintained using four loops, one for each corner of the matrix. The PrintSpiral method
in the following code snippet creates a matrix with elements arranged spirally in increasing
order. The method takes a parameter, n, and returns an integer array:*/
//PrintSpiral method

func PrintSpiral(n int) []int {

	var left int
	var top int
	var right int
	var bottom int

	left = 0
	top = 0
	right = n - 1
	bottom = n - 1
	var size int
	size = n * n
	var s []int
	s = make([]int, size)

	var i int
	i = 0
	for left < right {
		// work right, along top
		var c int
		for c = left; c <= right; c++ {
			s[top*n+c] = i
			i++
		}
		top++
		// work down right side
		var r int
		for r = top; r <= bottom; r++ {
			s[r*n+right] = i
			i++
		}
		right--
		if top == bottom {
			break
		}
		// work left, along bottom
		for c = right; c >= left; c-- {
			s[bottom*n+c] = i
			i++
		}
		bottom--
		// work up left side
		for r = bottom; r >= top; r-- {
			s[r*n+left] = i
			i++
		}
		left++
	}
	// center (last) element
	s[top*n+left] = i

	return s
}

/*In the following code snippet, the main method invokes the PrintSpiral method, which
takes the integer n and prints the integer values of the matrix spirally. The values returned
from the PrintSpiral method are printed as fields with a width of 2:*/

func main() {
	var n int
	n = 5
	var length int
	length = 2
	var i int
	var sketch int
	for i, sketch = range PrintSpiral(n) {
		fmt.Printf("%*d ", length, sketch)
		if i%n == n-1 {
			fmt.Println("")
		}
	}
}
