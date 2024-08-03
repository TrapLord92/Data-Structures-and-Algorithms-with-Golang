 example to compute the determinant of a 2x2 matrix:

```go
package main

import (
	"fmt"
)

// Function to calculate the determinant of a 2x2 matrix
func determinant(matrix [2][2]float64) float64 {
	return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
}

func main() {
	// Example matrix
	matrix := [2][2]float64{
		{1, 2},
		{3, 4},
	}

	// Calculate the determinant
	det := determinant(matrix)

	// Print the result
	fmt.Printf("The determinant of the matrix is: %.2f\n", det)
}
```

This code defines a function `determinant` to calculate the determinant of a 2x2 matrix. The `main` function provides an example matrix, computes its determinant using the `determinant` function, and prints the result.

When you run this code, you should see the output:

```
The determinant of the matrix is: -2.00
```