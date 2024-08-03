// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {
	/*A column matrix is an m x 1 matrix that has a single column of m elements. The following
	code snippet shows how to create a column matrix:*/
	var matrix = [4][1]int{{1}, {2}, {3}, {4}}

	fmt.Println(matrix)
}
