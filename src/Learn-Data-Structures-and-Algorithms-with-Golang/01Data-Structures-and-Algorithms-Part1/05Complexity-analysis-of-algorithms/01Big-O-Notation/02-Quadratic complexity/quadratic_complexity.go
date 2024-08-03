// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

/*An algorithm is of quadratic complexity if the processing time is proportional to the square
of the number of input elements. In the following case, the complexity of the algorithm is
10*10 = 100. The two loops have a maximum of 10. The quadratic complexity for a
multiplication table of n elements is O(n2).
Quadratic complexity, O(n2), is shown in the following example:*/
// main method
func main() {

	var k, l int

	for k = 1; k <= 10; k++ {
		fmt.Println(" Multiplication Table", k)
		for l = 1; l <= 10; l++ {
			var x int = l * k
			fmt.Println(x)
		}

	}
}
