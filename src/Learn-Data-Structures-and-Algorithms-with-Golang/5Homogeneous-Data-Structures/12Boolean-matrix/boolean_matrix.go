// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

/*Boolean matrix
A Boolean matrix is a matrix that consists of elements in the mth row and the nth column with
a value of 1. A matrix can be modified to be a Boolean matrix by making the values in
the mth row and the nth column equal to 1. In the following code, the Boolean matrix
transformation and print methods are shown in detail.*/

/* The changeMatrix method
transforms the input matrix in to a Boolean matrix by changing the row and column values
from 0 to 1 if the cell value is 1. The method takes the input matrix as the parameter and
returns the changed matrix, as shown in the following code:*/
//changeMatrix method
func changeMatrix(matrix [3][3]int) [3][3]int {
	var i int
	var j int
	var Rows [3]int
	var Columns [3]int

	var matrixChanged [3][3]int

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if matrix[i][j] == 1 {
				Rows[i] = 1
				Columns[j] = 1
			}

		}
	}

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if Rows[i] == 1 || Columns[j] == 1 {
				matrixChanged[i][j] = 1
			}

		}
	}

	return matrixChanged

}

// printMatrix method
/*The printMatrix method
In the following code snippet, the printMatrix method takes the input matrix and prints
the matrix values by row and traverses the columns for every row:*/
func printMatrix(matrix [3][3]int) {
	var i int
	var j int
	//var k int
	for i = 0; i < 3; i++ {

		for j = 0; j < 3; j++ {

			fmt.Printf("%d", matrix[i][j])

		}
		fmt.Printf("\n")
	}

}

// main method
/*The main method
The main method in the following code snippet invokes the changeMatrix method after
initializing the matrix. The changed matrix is printed after the invocation of the
changeMatrix method:*/
func main() {

	var matrix = [3][3]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	printMatrix(matrix)

	matrix = changeMatrix(matrix)

	printMatrix(matrix)

}
