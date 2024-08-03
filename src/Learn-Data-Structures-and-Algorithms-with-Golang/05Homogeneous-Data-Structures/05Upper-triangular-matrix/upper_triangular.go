// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {
	/*Upper triangular matrix
	An upper triangular matrix consists of elements with a value of zero below the main
	diagonal. The following code creates an upper triangular matrix:*/

	var matrix = [3][3]int{
		{1, 2, 3},
		{0, 1, 4},
		{0, 0, 1}}

	fmt.Println(matrix)
}
