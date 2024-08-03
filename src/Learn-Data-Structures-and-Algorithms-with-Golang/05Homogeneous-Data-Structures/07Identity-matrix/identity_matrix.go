// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// identity method
/*Identity matrix
An identity matrix is a unit matrix with ones are on the main diagonal and zeros are
elsewhere. The following code snippet creates an identity matrix:*/
func Identity(order int) [][]float64 {
	var matrix [][]float64
	matrix = make([][]float64, order)
	var i int
	for i = 0; i < order; i++ {
		var temp []float64
		temp = make([]float64, order)
		temp[i] = 1
		matrix[i] = temp
	}
	return matrix
}

// main method
func main() {
	fmt.Println(Identity(4))
}
