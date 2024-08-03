// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {
	/*Lower triangular matrix
	  A lower triangular matrix consists of elements that have a value of zero above the main
	  diagonal. The following code snippet shows how to create a lower triangular matrix:*/
	var matrix = [3][3]int{
		{1, 0, 0},
		{1, 1, 0},
		{2, 1, 1}}

	fmt.Println(matrix)
}
