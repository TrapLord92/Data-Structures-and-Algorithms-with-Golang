// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {
	/*Null matrix
	A null or a zero matrix is a matrix entirely consisting of zero values, as shown in the
	following code snippet:*/

	var matrix = [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0}}

	fmt.Println(matrix)
}
