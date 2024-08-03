// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

/*
In this section, we will look at the basic operations on the two-dimensional matrix. Let's
start with initializing the matrices.
matrix1 and matrix2 are initialized in the following code snippet:
*/
var matrix1 = [2][2]int{
	{4, 5},
	{1, 2},
}
var matrix2 = [2][2]int{
	{6, 7},
	{3, 4},
}

/*The add, subtract, multiply, transpose, and inversion operations are presented here*/

// add method
/*The add method
The add method adds the elements of two 2 x 2 matrices. The following code returns the
created matrix by adding the two matrices:*/
func add(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var m int
	var l int
	var sum [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			sum[l][m] = matrix1[l][m] + matrix2[l][m]
		}
	}
	return sum
}

// subtract method
/*The subtract method
The subtract method subtracts the elements of two 2 x 2 matrices. The subtract method
in the following snippet subtracts the elements of matrix1 and matrix2. This method
returns the resulting matrix after subtraction:*/
func subtract(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var m int
	var l int
	var difference [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			difference[l][m] = matrix1[l][m] - matrix2[l][m]
		}
	}
	return difference

}

// multiply method
/*The multiply method
The multiply method multiplies the elements of two 2 x 2 matrices. The multiplication of
two matrices, matrix1 and matrix2, is shown in the following snippet. The matrix
that's generated after the multiplication is returned by the multiply method:*/
func multiply(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var m int
	var l int
	var n int
	var product [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			var productSum int = 0
			for n = 0; n < 2; n++ {
				productSum = productSum + matrix1[l][n]*matrix2[n][m]

			}
			product[l][m] = productSum
		}
	}
	return product
}

// transpose method
/*The transpose method
The transpose of a matrix is achieved using the transpose method. This method takes the
matrix as a parameter and returns the transposed matrix:*/

/*
The transpose of a matrix is an operation that flips a matrix over its diagonal, switching the row and column indices of each element.

For example, given a matrix \( A \):

\[
A = \begin{bmatrix}
1 & 2 \\
3 & 4 \\
\end{bmatrix}
\]

The transpose of \( A \), denoted as \( A^T \), is:

\[
A^T = \begin{bmatrix}
1 & 3 \\
2 & 4 \\
\end{bmatrix}
\]
*/
func transpose(matrix1 [2][2]int) [2][2]int {
	var m int
	var l int
	var transMatrix [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			transMatrix[l][m] = matrix1[m][l]
		}
	}
	return transMatrix
}

// determinant method
/*The determinant method
The determinant method calculates the determinant of the matrix. The determinant
method in the following code snippet calculates the determinant value of a matrix. The
method takes the matrix and returns a float32 value, which is the determinant of the
matrix*/
/*The determinant of a matrix is a scalar value that is computed from the elements of a square matrix. It provides important properties about the matrix, such as whether the matrix is invertible and the volume scaling factor of the linear transformation represented by the matrix.

For a 2x2 matrix \( A \):

\[
A = \begin{bmatrix}
a & b \\
c & d \\
\end{bmatrix}
\]

The determinant, denoted as \(\det(A)\) or \(|A|\), is calculated as:

\[
\det(A) = ad - bc
\]

For example, given:

\[
A = \begin{bmatrix}
3 & 8 \\
4 & 6 \\
\end{bmatrix}
\]

The determinant of \( A \) is:

\[
\det(A) = (3 \times 6) - (8 \times 4) = 18 - 32 = -14
\]*/
func determinant(matrix1 [2][2]int) float64 {

	var det float64

	det = det + float64(((matrix1[0][0] * matrix1[1][1]) - (matrix1[0][1] * matrix1[1][0])))

	return det
}

// inverse method
/*The inverse method
The inverse method returns the inverse of the matrix, which is passed as a parameter.
This is shown in the following snippet:*/
func inverse(matrix [2][2]int) [][]float64 {
	var det float64

	det = determinant(matrix)

	var invmatrix [2][2]float64
	invmatrix[0][0] = float64(matrix[1][1]) / det
	invmatrix[0][1] = -1 * float64(matrix[0][1]) / det
	invmatrix[1][0] = -1 * float64(matrix[1][0]) / det
	invmatrix[1][1] = float64(matrix[0][0]) / det

	return [][]float64{
		{invmatrix[0][0], invmatrix[0][1]},
		{invmatrix[1][0], invmatrix[1][1]},
	}
}

// main method
func main() {

	var matrix1 = [2][2]int{
		{4, 5},
		{1, 2}}
	var matrix2 = [2][2]int{
		{6, 7},
		{3, 4}}
	/*The sum between the two matrices is the result of calling the add method. The parameters
	that are passed are the matrices to be added, as shown here:*/
	var sum [2][2]int
	sum = add(matrix1, matrix2)

	fmt.Println(sum)
	/*The difference between two matrices is the result of calling the subtract method. The
	  parameters that are passed are the matrices to be subtracted, as shown here:*/
	var difference [2][2]int
	difference = subtract(matrix1, matrix2)

	fmt.Println(difference)
	/*The product of two matrices is calculated using the multiply method in the following
	code snippet, which takes the two matrices as parameters:*/

	var product [2][2]int
	product = multiply(matrix1, matrix2)

	fmt.Println(product)

}
