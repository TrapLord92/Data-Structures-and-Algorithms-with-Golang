// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {
	/*A row matrix is a 1 x m matrix consisting of a single row of m elements, as shown here:*/

	var matrix = [1][3]int{{1, 2, 3}}

	fmt.Println(matrix)
}
